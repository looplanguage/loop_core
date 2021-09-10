package object

import "fmt"

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() string    { return BOOLEAN }
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }
