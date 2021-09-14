package object

import (
	"fmt"
	"strings"
)

type HashKey struct {
	Type  string
	Value uint64
}

type Hashable interface {
	Hash() HashKey
}

type HashPair struct {
	Key   Object
	Value Object
}

type Hash interface {
	Hash() HashKey
}

type HashMap struct {
	Pairs map[HashKey]HashPair
}

func (h *HashMap) Type() string {
	return HASHMAP
}

func (h *HashMap) Inspect() string {
	values := []string{}

	for _, val := range h.Pairs {
		values = append(values, fmt.Sprintf("%v: %v", val.Key.Inspect(), val.Value.Inspect()))
	}

	return fmt.Sprintf("{%v}", strings.Join(values, ", "))
}
