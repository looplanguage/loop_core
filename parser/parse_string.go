package parser

import "git.kanersps.pw/loop/models/ast"

func (p *Parser) parseString() ast.Expression {
	return &ast.String{Value: p.CurrentToken.Literal}
}
