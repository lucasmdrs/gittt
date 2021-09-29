package gittt

type trigger func(g *Gittt, data []byte) error

var availableTriggers = map[string]trigger{
	PREvent:      PullRequestTrigger,
	IssuesEvent:  IssuesTrigger,
	ReleaseEvent: ReleaseTrigger,
}
