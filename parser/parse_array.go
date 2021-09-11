package parser

import (
	"git.kanersps.pw/loop/models/ast"
	"git.kanersps.pw/loop/models/tokens"
	"git.kanersps.pw/loop/parser/precedence"
)

func (p *Parser) parseArray() ast.Expression {
	array := &ast.Array{Token: p.CurrentToken}

	p.NextToken()

	for p.CurrentToken.Type != tokens.RightBracket {
		if p.CurrentToken.Type == tokens.Comma {
			p.NextToken()
		}

		exp := p.parseExpression(precedence.LOWEST)

		if exp != nil {
			array.Elements = append(array.Elements, exp)
		}

		p.NextToken()
	}

	p.NextToken()

	return array
}
