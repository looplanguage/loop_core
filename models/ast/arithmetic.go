package ast

import (
	"fmt"
	"git.kanersps.pw/loop/models/tokens"
)

type IntegerLiteral struct {
	Token tokens.Token
	Value int64
}

func (i *IntegerLiteral) expressionNode()      {}
func (i *IntegerLiteral) TokenLiteral() string { return i.Token.Name() }
func (i *IntegerLiteral) String() string       { return i.Token.Literal }

type SuffixExpression struct {
	Token    tokens.Token
	Left     Expression
	Right    Expression
	Operator string
}

func (se *SuffixExpression) expressionNode()      {}
func (se *SuffixExpression) TokenLiteral() string { return se.Token.Name() }
func (se *SuffixExpression) String() string {
	return fmt.Sprintf("(%s %s %s)", se.Left.String(), se.Operator, se.Right.String())
}
