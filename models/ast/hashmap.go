package ast

import (
	"fmt"
	"git.kanersps.pw/loop/models/tokens"
	"strings"
)

type Hashmap struct {
	Token  tokens.Token
	Values map[Expression]Expression
}

func (h *Hashmap) expressionNode()      {}
func (h *Hashmap) TokenLiteral() string { return h.Token.Literal }
func (h *Hashmap) String() string {
	value := "{"

	values := []string{}

	for key, val := range h.Values {
		values = append(values, fmt.Sprintf("%v: %v", key.String(), val.String()))
	}

	value += strings.Join(values, ", ") + "}"

	return value
}
