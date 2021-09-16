package ast

import (
	"fmt"
	"github.com/looplanguage/loop/models/tokens"
)

type Return struct {
	Token tokens.Token
	Value Expression
}

func (r *Return) statementNode()       {}
func (r *Return) TokenLiteral() string { return r.Token.Literal }
func (r *Return) String() string       { return fmt.Sprintf("return %s", r.Value.String()) }
