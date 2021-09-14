package object

type Null struct{}

func (n *Null) Type() string    { return NONE }
func (n *Null) Inspect() string { return "null" }
