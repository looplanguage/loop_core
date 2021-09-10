package ast

import (
	"fmt"
	"git.kanersps.pw/loop/models/tokens"
)

type Identifier struct {
	Token tokens.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) statementNode()       {}
func (i *Identifier) TokenLiteral() string { return i.Token.Name() }
func (i *Identifier) String() string {
	return fmt.Sprintf(i.Value)
}

type VariableDeclaration struct {
	Token      tokens.Token
	Identifier *Identifier
	Value      Expression
}

func (vd *VariableDeclaration) statementNode()       {}
func (vd *VariableDeclaration) TokenLiteral() string { return vd.Token.Name() }
func (vd *VariableDeclaration) String() string {
	return fmt.Sprintf("var %v = %v", vd.Identifier.Value, vd.Value.String())
}

type Boolean struct {
	Token tokens.Token
	Value bool
}

func (b *Boolean) expressionNode()      {}
func (b *Boolean) TokenLiteral() string { return b.Token.Name() }
func (b *Boolean) String() string {
	return fmt.Sprintf("%t", b.Value)
}
