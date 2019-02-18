package gittt

const (
	ghEventHeader = "X-GitHub-Event"
	AnyEvent      = "*"
	PREvent       = "pull_request"
	IssuesEvent   = "issues"
	ReleaseEvent  = "release"
)

type trigger func(g *gittt, data []byte) error

var availableTriggers = map[string]trigger{
	PREvent:      PullRequestTrigger,
	IssuesEvent:  IssuesTrigger,
	ReleaseEvent: ReleaseTrigger,
}
