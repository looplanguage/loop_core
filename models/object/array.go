package object

import (
	"fmt"
	"strings"
)

type Array struct {
	Elements []Object
}

func (a *Array) Type() string { return ARRAY }
func (a *Array) Inspect() string {
	value := "["

	var values []string

	for _, el := range a.Elements {
		values = append(values, el.Inspect())
	}

	value += fmt.Sprintf("%s]", strings.Join(values, ", "))

	return value
}
