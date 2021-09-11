package object

import "fmt"

type String struct {
	Value string
}

func (s *String) Type() string    { return STRING }
func (s *String) Inspect() string { return fmt.Sprintf("%q", s.Value) }
