package gittt

import (
	"encoding/json"
	"log"
)

type PullRequest struct {
	State  string          `json:"action"`
	PRInfo PullRequestInfo `json:"pull_request"`
}

type PullRequestInfo struct {
	Title      string `json:"title"`
	FromBranch Branch `json:"head"`
	IntoBranch Branch `json:"base"`
	Merged     bool   `json:"merged"`
}

func PullRequestTrigger(g *Gittt, data []byte) error {
	var pr PullRequest
	err := json.Unmarshal(data, &pr)
	if err != nil {
		return err
	}

	log.Println("Pull Request Event Received")
	actions := g.matchConditionals(PREvent, pr)
	for _, action := range actions {
		action.Do(pr)
	}

	return nil
}

func (g *Gittt) ConditionPRMergedInAnyOf(data interface{}, branches ...interface{}) bool {
	if pr, ok := data.(PullRequest); ok {
		if pr.State == "closed" {
			for _, branch := range branches {
				if pr.PRInfo.IntoBranch.Ref == branch.(string) {
					log.Printf("Condition match: PR merged on %s\n", branch.(string))
					return true
				}
			}
		}
	}
	return false
}
