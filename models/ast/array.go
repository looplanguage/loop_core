package ast

import (
	"fmt"
	"github.com/looplanguage/loop/models/tokens"
	"strings"
)

type IndexExpression struct {
	Token tokens.Token
	Index Expression
	Value Expression
}

func (ie *IndexExpression) expressionNode()      {}
func (ie *IndexExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IndexExpression) String() string {
	return fmt.Sprintf("%s[%s]", ie.Value.TokenLiteral(), ie.Index.TokenLiteral())
}

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
