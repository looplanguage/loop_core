package ast

import "git.kanersps.pw/loop/models/tokens"

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// The entire program will be contained in here
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if p.Statements[0] != nil {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	statements := ""

	for _, stmt := range p.Statements {
		statements = statements + stmt.String()
	}

	return statements
}

type ExpressionStatement struct {
	Token tokens.Token
	Expression Expression
}

func (e *ExpressionStatement) statementNode() { }
func (e *ExpressionStatement) TokenLiteral() string { return e.Token.Name() }
func (e *ExpressionStatement) String() string {
	if e.Expression != nil {
		return e.Expression.String()
	}

	return ""
}
