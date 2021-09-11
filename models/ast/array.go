package ast

import (
	"fmt"
	"git.kanersps.pw/loop/models/tokens"
	"strings"
)

type Array struct {
	Token    tokens.Token
	Elements []Expression
}

func (a *Array) expressionNode()      {}
func (a *Array) TokenLiteral() string { return a.Token.Literal }
func (a *Array) String() string {
	value := "["

	var values []string

	for _, el := range a.Elements {
		values = append(values, el.String())
	}

	value += fmt.Sprintf("%s]", strings.Join(values, ", "))

	return value
}
