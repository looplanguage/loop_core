package parser

import (
	"git.kanersps.pw/loop/models/ast"
	"git.kanersps.pw/loop/models/tokens"
)

func (p *Parser) parseStatement() ast.Statement {
	switch p.CurrentToken.Type {
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	expression := &ast.ExpressionStatement{
		Token: p.CurrentToken,
	}

	expression.Expression = p.parseExpression()

	return expression
}

func (p *Parser) parseExpression() ast.Expression {
	prefixFn := p.prefixParsers[p.CurrentToken.Type]

	if prefixFn == nil {
		// TODO: add parser error
		return nil
	}

	expression := prefixFn()

	for !p.peekTokenIs(tokens.Semicolon) {
		suffixFn := p.suffixParsers[p.PeekToken.Type]

		if suffixFn == nil {
			return expression
		}

		p.NextToken()

		expression = suffixFn(expression)
	}

	return expression
}
