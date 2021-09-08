package object

import "fmt"

const (
	INTEGER = "INTEGER"
)

type Object interface {
	Type() string
	Inspect() string
}

type Integer struct {
	Value int64
}

func (i *Integer) Type() string    { return INTEGER }
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }
