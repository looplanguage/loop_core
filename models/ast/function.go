package ast

import (
	"fmt"
	"git.kanersps.pw/loop/models/tokens"
	"strings"
)

type CallExpression struct {
	Token      tokens.Token
	Function   Expression
	Parameters []Expression
}

func (c *CallExpression) expressionNode()      {}
func (c *CallExpression) TokenLiteral() string { return c.Token.Literal }
func (c *CallExpression) String() string {
	parameters := []string{}

	for _, p := range c.Parameters {
		parameters = append(parameters, p.String())
	}

	return fmt.Sprintf("%v(%v)", c.Token.Literal, strings.Join(parameters, ", "))
}

type Function struct {
	Token      tokens.Token
	Parameters []*Identifier
	Body       BlockStatement
}

func (f *Function) expressionNode()      {}
func (f *Function) TokenLiteral() string { return f.Token.Literal }
func (f *Function) String() string {
	value := "func("
	parameters := []string{}

	for _, param := range f.Parameters {
		parameters = append(parameters, param.Token.Literal)
	}

	value += strings.Join(parameters, ", ")

	value += ") { "

	value += f.String()

	value += " }"

	return value
}
