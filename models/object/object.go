package object

import "fmt"

const (
	INTEGER = "INTEGER"
	ERROR   = "ERROR"
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

type Error struct {
	Message string
	Line    int64
	Column  int64
}

func (e *Error) Type() string    { return ERROR }
func (e *Error) Inspect() string { return fmt.Sprintf("Exception on line %d: %s", e.Line, e.Message) }
