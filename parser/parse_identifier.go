package parser

import (
	"github.com/looplanguage/loop/models/ast"
)

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{
		Token: p.CurrentToken,
		Value: p.CurrentToken.Literal,
	}
}
