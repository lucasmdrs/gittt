package gittt

import "log"

type action struct {
	args       []interface{}
	actionFunc func(data interface{}, args ...interface{})
}

func (a *action) Do(data interface{}) {
	a.actionFunc(data, a.args...)
}

func (g *Gittt) ActionLogPayload(data interface{}, args ...interface{}) {
	log.Printf("%+v\n", data)
}
