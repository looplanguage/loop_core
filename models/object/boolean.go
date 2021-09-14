package object

import "fmt"

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() string    { return BOOLEAN }
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }
func (b *Boolean) Hash() HashKey {
	value := 0

	if b.Value {
		value = 1
	}

	return HashKey{
		Type:  b.Type(),
		Value: uint64(value),
	}
}
