package parser

import (
	"git.kanersps.pw/loop/models/ast"
)

func (p *Parser) parseBoolean() ast.Expression {
	return &ast.Boolean{
		Token: p.CurrentToken,
		Value: p.CurrentToken.Literal == "true",
	}
}
