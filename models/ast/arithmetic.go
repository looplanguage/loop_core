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

type ConditionalStatement struct {
	Token         tokens.Token
	Condition     Expression
	ElseCondition *BlockStatement
	ElseStatement *ConditionalStatement
	Body          *BlockStatement
}

func (cs *ConditionalStatement) expressionNode()      {}
func (cs *ConditionalStatement) TokenLiteral() string { return cs.Token.Literal }
func (cs *ConditionalStatement) String() string {
	value := fmt.Sprintf("if(%s) { \n", cs.Condition.String())

	for _, stmt := range cs.Body.Statements {
		value += "\t" + stmt.String() + "\n"
	}

	value += "} "

	if cs.ElseCondition != nil {
		if len(cs.ElseCondition.Statements) > 0 {
			value += " else { "

			for _, stmt := range cs.ElseCondition.Statements {
				value += "\t" + stmt.String() + "\n"
			}

			value += "}"
		}
	}

	return value
}
