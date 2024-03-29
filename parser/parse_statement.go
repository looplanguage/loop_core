package parser

import (
	"fmt"
	"github.com/looplanguage/loop/models/ast"
	"github.com/looplanguage/loop/models/tokens"
	"github.com/looplanguage/loop/parser/precedence"
)

func (p *Parser) parseStatement() ast.Statement {
	switch p.CurrentToken.Type {

	case tokens.Return:
		return p.parseReturn()
	case tokens.VariableDeclaration:
		return p.parseVariable()
	case tokens.Import:
		return p.parseImportStatement()
	case tokens.Export:
		return p.parseExportStatement()
	case tokens.Identifier:
		if p.peekTokenIs(tokens.Assign) {
			return p.parseAssignExpression()
		}

		return p.parseExpressionStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseExpressionStatement() ast.Statement {
	expression := &ast.ExpressionStatement{
		Token: p.CurrentToken,
	}

	expression.Expression = p.parseExpression(precedence.LOWEST)

	if p.PeekToken.Type == tokens.Semicolon {
		p.NextToken()
	}

	if arr, ok := expression.Expression.(*ast.IndexExpression); ok {
		fmt.Println(":)")
		if p.peekTokenIs(tokens.Assign) {
			p.NextToken()
			p.NextToken()

			value := p.parseExpression(precedence.LOWEST)

			return &ast.IndexAssign{
				Token:  p.CurrentToken,
				Object: arr.Value,
				Index:  arr.Index,
				Value:  value,
			}
		}
	}

	return expression
}

func (p *Parser) parseExpression(pre int) ast.Expression {
	prefixFn := p.prefixParsers[p.CurrentToken.Type]

	if prefixFn == nil {
		p.AddError(fmt.Sprintf("no prefix parser for %q. expected=prefixParseFn. got=nil", p.CurrentToken.Name()))
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

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.NextToken()
	exp := p.parseExpression(precedence.LOWEST)

	p.NextToken()
	if p.CurrentToken.Type != tokens.RightParenthesis {
		p.AddError(fmt.Sprintf("wrong token. expected=%q. got=%q", ")", p.CurrentToken.Literal))
		return nil
	}

	return exp
}
