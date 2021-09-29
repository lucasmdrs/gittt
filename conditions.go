package gittt

type condition struct {
	event    EventType
	arg      []interface{}
	evalFunc func(interface{}, ...interface{}) bool
	actions  []action
}

func (c *condition) AddAction(a action) error {
	c.actions = append(c.actions, a)
	return nil
}

func (g *Gittt) ConditionAlways(data interface{}, args ...interface{}) bool {
	return true
}
