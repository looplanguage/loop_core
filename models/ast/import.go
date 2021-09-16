package ast

import (
	"fmt"
	"github.com/looplanguage/loop/models/tokens"
)

type Export struct {
	Token      tokens.Token
	Expression Expression
}

func (e *Export) statementNode()       {}
func (e *Export) TokenLiteral() string { return e.Token.Literal }
func (e *Export) String() string {
	return fmt.Sprintf("export %s", e.Expression.String())
}

type Import struct {
	Token      tokens.Token
	File       string
	Identifier string
	Caller     string
}

func (i *Import) statementNode()       {}
func (i *Import) TokenLiteral() string { return i.Token.Literal }
func (i *Import) String() string {
	return fmt.Sprintf("import %q as %s", i.File, i.Identifier)
}
