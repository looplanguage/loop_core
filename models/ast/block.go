package ast

import "git.kanersps.pw/loop/models/tokens"

type BlockStatement struct {
	Token      tokens.Token
	Statements []Statement
}

func (b *BlockStatement) statementNode()       {}
func (b *BlockStatement) TokenLiteral() string { return b.Token.Literal }
func (b *BlockStatement) String() string {
	value := ""

	for _, v := range b.Statements {
		value += v.String() + "\n"
	}

	return value
}
