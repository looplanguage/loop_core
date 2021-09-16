package ast

import (
	"fmt"
	"github.com/looplanguage/loop/models/tokens"
)

type While struct {
	Token     tokens.Token
	Condition Expression
	Block     *BlockStatement
}

func (w *While) expressionNode()      {}
func (w *While) TokenLiteral() string { return w.Token.Literal }
func (w *While) String() string {
	value := fmt.Sprintf("while(%v) {\n", w.Condition.String())

	for _, stmt := range w.Block.Statements {
		value += "\t" + stmt.String() + "\n"
	}

	value += "}"

	return value
}
