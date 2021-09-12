package object

type None struct{}

func (n *None) Type() string    { return NONE }
func (n *None) Inspect() string { return "undefined" }
