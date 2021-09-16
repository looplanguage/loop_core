package parser

import "github.com/looplanguage/loop/models/ast"

func (p *Parser) parseString() ast.Expression {
	return &ast.String{Token: p.CurrentToken, Value: p.CurrentToken.Literal}
}
