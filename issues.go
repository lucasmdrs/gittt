package gittt

import (
	"encoding/json"
	"log"
)

type Issue struct {
	Info IssueInfo
}

type IssueInfo struct {
	State  string  `json:"state"`
	Title  string  `json:"title"`
	Body   string  `json:"body"`
	Labels []Label `json:"labels"`
}

func IssuesTrigger(g *Gittt, data []byte) error {
	var i Issue
	err := json.Unmarshal(data, &i)
	if err != nil {
		return err
	}

	log.Println("Issue Event Received")
	actions := g.matchConditionals(i)
	for _, action := range actions {
		action.Do(i)
	}

	return nil
}

func (g *Gittt) IssueLabelsIsOneOf(data interface{}, labels ...interface{}) bool {
	if i, ok := data.(Issue); ok {
		for _, label := range labels {
			for _, l := range i.Info.Labels {
				if label.(string) == l.Name {
					log.Printf("Condition match: labels contains %s\n", label.(string))
					return true
				}
			}
		}
	}
	return false
}

func (g *Gittt) IssueIsClosed(data interface{}, args ...interface{}) bool {
	if i, ok := data.(Issue); ok {
		if i.Info.State == "closed" {
			log.Println("Condition match: Issue is closed.")
			return true
		}
	}
	return false
}
