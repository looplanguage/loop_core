package parser

import (
	"git.kanersps.pw/loop/models/ast"
	"git.kanersps.pw/loop/models/tokens"
	"git.kanersps.pw/loop/parser/precedence"
)

func (p *Parser) parseStatement() ast.Statement {
	switch p.CurrentToken.Type {
	case tokens.VariableDeclaration:
		return p.parseVariable()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	expression := &ast.ExpressionStatement{
		Token: p.CurrentToken,
	}

	expression.Expression = p.parseExpression(precedence.LOWEST)

	return expression
}

func (p *Parser) parseExpression(pre int) ast.Expression {
	prefixFn := p.prefixParsers[p.CurrentToken.Type]

	if prefixFn == nil {
		// TODO: add parser error
		return nil
	}

	expression := prefixFn()

	for !p.peekTokenIs(tokens.Semicolon) && pre < p.peekPrecedence() {
		suffixFn := p.suffixParsers[p.PeekToken.Type]

		if suffixFn == nil {
			return expression
		}

		p.NextToken()

		expression = suffixFn(expression)
	}

	return expression
}
