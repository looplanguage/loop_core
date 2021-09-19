package object

import "fmt"

const (
	INTEGER           = "INTEGER"
	ERROR             = "ERROR"
	FUNCTION          = "FUNCTION"
	BOOLEAN           = "BOOLEAN"
	ARRAY             = "ARRAY"
	RETURN            = "RETURN"
	BUILTIN           = "BUILTIN"
	STRING            = "STRING"
	NONE              = "NONE"
	HASHMAP           = "HASHMAP"
	EXPORT            = "EXPORT"
	IMPORT            = "IMPORT"
	COMPILED_FUNCTION = "COMPILED_FUNCTION"
	CLOSURE           = "CLOSURE"
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
func (i *Integer) Hash() HashKey {
	return HashKey{
		Type:  i.Type(),
		Value: uint64(i.Value),
	}
}

type Error struct {
	Message string
	Line    int
	Column  int
}

func (e *Error) Type() string { return ERROR }
func (e *Error) Inspect() string {
	return fmt.Sprintf("EvaluationException at %d,%d: %s", e.Line, e.Column, e.Message)
}
