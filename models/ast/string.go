package ast

import (
	"fmt"
	"github.com/looplanguage/loop/models/tokens"
)

type String struct {
	Token tokens.Token
	Value string
}

func (s *String) expressionNode()      {}
func (s *String) TokenLiteral() string { return s.Token.Literal }
func (s *String) String() string       { return fmt.Sprintf("%q", s.Value) }
