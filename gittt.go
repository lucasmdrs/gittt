package gittt

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const ghEventHeader = "X-GitHub-Event"

type Gittt struct {
	triggers   map[string]trigger
	conditions []condition
}

func Init() *Gittt {
	return &Gittt{
		triggers:   make(map[string]trigger),
		conditions: make([]condition, 0),
	}
}

func (g *Gittt) Handler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	event := r.Header.Get(ghEventHeader)
	t, enabled := g.triggers[event]
	if !enabled {
		log.Printf("Invalid Event: %s\n", event)
		http.Error(w, "Invalid Event", http.StatusBadRequest)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Invalid payload")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	go t(g, data)
}

func (g *Gittt) ListenForEvents(eventTypes ...string) error {
	currentTriggers := g.triggers
	for _, event := range eventTypes {
		if t, exist := availableTriggers[event]; exist {
			g.triggers[event] = t
			continue
		}
		g.triggers = currentTriggers
		return fmt.Errorf(`event "%s" is not available`, event)
	}
	return nil
}

func (g *Gittt) ListenAllEvents() {
	for event, t := range availableTriggers {
		g.triggers[event] = t
	}
}

func (g *Gittt) ConditionBuilder(onEvent EventType, conditionFunc func(data interface{}, args ...interface{}) bool, args ...interface{}) condition {
	return condition{
		event:    onEvent,
		arg:      args,
		evalFunc: conditionFunc,
	}
}

func (g *Gittt) ActionBuilder(actionFunc func(data interface{}, args ...interface{}), args ...interface{}) action {
	return action{
		args:       args,
		actionFunc: actionFunc,
	}
}

func (g *Gittt) AddConditions(conditions ...condition) {
	g.conditions = append(g.conditions, conditions...)
}
func (g *Gittt) matchConditionals(event EventType, data interface{}) (actions []action) {
	for _, c := range g.conditions {
		if (c.event == event || c.event == AnyEvent) && c.evalFunc(data, c.arg...) {
			actions = append(actions, c.actions...)
		}
	}
	return actions
}
