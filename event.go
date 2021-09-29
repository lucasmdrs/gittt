package gittt

type EventType string

const (
	AnyEvent     = "*"
	PREvent      = "pull_request"
	IssuesEvent  = "issues"
	ReleaseEvent = "release"
)
