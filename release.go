package gittt

import (
	"encoding/json"
	"log"
)

type Release struct {
	Action      string      `json:"action"`
	ReleaseInfo ReleaseInfo `json:"release"`
}

type ReleaseInfo struct {
	Name         string `json:"name"`
	Draft        bool   `json:"draft"`
	Tag          string `json:"tag_name"`
	BranchTarget string `json:"target_commitish"`
}

func ReleaseTrigger(g *Gittt, data []byte) error {
	var r Release
	err := json.Unmarshal(data, &r)
	if err != nil {
		return err
	}

	log.Println("Release Event Received")
	actions := g.matchConditionals(r)
	for _, action := range actions {
		action.Do(r)
	}

	return nil
}

func (g *Gittt) ConditionReleaseFromOneOf(data interface{}, branches ...interface{}) bool {
	if r, ok := data.(Release); ok {
		if !r.ReleaseInfo.Draft {
			for _, branch := range branches {
				if r.ReleaseInfo.BranchTarget == branch.(string) {
					log.Printf("Condition match: Release created from %s\n", branch.(string))
					return true
				}
			}

		}
	}
	return false
}
