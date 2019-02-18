package gittt

import (
	"encoding/json"
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

func IssuesTrigger(g *gittt, data []byte) error {
	var i Issue
	err := json.Unmarshal(data, &i)
	if err != nil {
		return err
	}

	actions := g.matchConditionals(i)
	for _, action := range actions {
		action.Do(i)
	}

	return nil
}

func (g *gittt) IssueLabelsIsOneOf(data interface{}, labels ...interface{}) bool {
	if i, ok := data.(Issue); ok {
		for _, label := range labels {
			for _, l := range i.Info.Labels {
				if label.(string) == l.Name {
					return true
				}
			}
		}
	}
	return false
}

func (g *gittt) IssueIsClosed(data interface{}, args ...interface{}) bool {
	if i, ok := data.(Issue); ok {
		return i.Info.State == "closed"
	}
	return false
}
