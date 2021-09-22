package ast

import (
	"fmt"
	"github.com/looplanguage/loop/models/tokens"
)

type Assign struct {
	Token      tokens.Token
	Identifier *Identifier
	Value      Expression
}

func (vd *Assign) statementNode()       {}
func (vd *Assign) TokenLiteral() string { return vd.Token.Name() }
func (vd *Assign) String() string {
	return fmt.Sprintf("%v = %v", vd.Identifier.Value, vd.Value.String())
}
