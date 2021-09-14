package object

import (
	"fmt"
	"hash/fnv"
)

type String struct {
	Value string
}

func (s *String) Type() string    { return STRING }
func (s *String) Inspect() string { return fmt.Sprintf("%q", s.Value) }
func (s *String) Hash() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))

	return HashKey{
		Type:  s.Type(),
		Value: h.Sum64(),
	}
}
