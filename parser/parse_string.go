package parser

import "git.kanersps.pw/loop/models/ast"

func (p *Parser) parseString() ast.Expression {
	return &ast.String{Token: p.CurrentToken, Value: p.CurrentToken.Literal}
}
