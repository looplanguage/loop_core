package parser

import (
	"git.kanersps.pw/loop/models/ast"
)

func (p *Parser) parseStatement() ast.Statement {
	switch p.CurrentToken.Type {
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	expression := &ast.ExpressionStatement{
		Token:      p.CurrentToken,
	}

	expression.Expression = p.parseExpression()

	return expression
}

func (p *Parser) parseExpression() ast.Expression {
	if fn, ok := p.prefixParsers[p.CurrentToken.Type]; ok {
		result := fn()

		return result
	}

	return nil
}
