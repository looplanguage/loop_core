package ast

import (
	"git.kanersps.pw/loop/models/tokens"
)

type IntegerLiteral struct {
	Token tokens.Token
	Value int64
}

func (i *IntegerLiteral) expressionNode()      {}
func (i *IntegerLiteral) TokenLiteral() string { return i.Token.Name() }
func (i *IntegerLiteral) String() string       { return i.Token.Literal }
