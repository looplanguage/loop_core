package parser

import (
	"fmt"
	"github.com/looplanguage/loop/models/ast"
	"github.com/looplanguage/loop/models/tokens"
	"github.com/looplanguage/loop/parser/precedence"
)

func (p *Parser) parseArray() ast.Expression {
	array := &ast.Array{Token: p.CurrentToken}

	p.NextToken()

	for p.CurrentToken.Type != tokens.RightBracket && p.CurrentToken.Type != tokens.EOF {
		exp := p.parseExpression(precedence.LOWEST)

		if exp != nil {
			array.Elements = append(array.Elements, exp)
		}

		p.NextToken()

		if p.CurrentToken.Type == tokens.Comma {
			p.NextToken()
		}
	}

	if p.CurrentToken.Type == tokens.Semicolon {
		p.NextToken()
	}

	return array
}

func (p *Parser) parseIndexExpression(left ast.Expression) ast.Expression {
	indexExpression := &ast.IndexExpression{
		Token: p.CurrentToken,
		Value: left,
	}

	p.NextToken()

	exp := p.parseExpression(precedence.LOWEST)

	if exp != nil {
		indexExpression.Index = exp
	}

	p.NextToken()

	if p.CurrentToken.Type != tokens.RightBracket {
		p.AddError(fmt.Sprintf("wrong token. expected=%q. got=%q", "]", p.CurrentToken.Literal))
		return nil
	}

	return indexExpression
}
