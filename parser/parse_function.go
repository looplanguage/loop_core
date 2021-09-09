package parser

import (
	"fmt"
	"git.kanersps.pw/loop/models/ast"
	"git.kanersps.pw/loop/models/tokens"
	"git.kanersps.pw/loop/parser/precedence"
)

func (p *Parser) parseCallExpression(fn ast.Expression) ast.Expression {
	return &ast.CallExpression{
		Token:      p.CurrentToken,
		Function:   fn,
		Parameters: p.parseCallArguments(),
	}
}

/*
func (p *Parser) parseCallArguments() []ast.Expression {
	args := []ast.Expression{}
	if p.peekTokenIs(tokens.RightParentheses) {
		p.ExtractToken()
		return args
	}
	p.ExtractToken()
	args = append(args, p.parseExpression(LOWEST))
	for p.peekTokenIs(tokens.Comma) {
		p.ExtractToken()
		p.ExtractToken()
		args = append(args, p.parseExpression(LOWEST))
	}
	if !p.expectPeek(tokens.RightParentheses) {
		return nil
	}
	return args
}
*/

func (p *Parser) parseCallArguments() []ast.Expression {
	var arguments []ast.Expression

	if p.PeekToken.Type == tokens.RightParenthesis {
		p.NextToken()
		return arguments
	}

	p.NextToken()

	arguments = append(arguments, p.parseExpression(precedence.LOWEST))

	for p.PeekToken.Type == tokens.Comma {
		p.NextToken() // Go to comma
		p.NextToken() // Go to next arg

		arguments = append(arguments, p.parseExpression(precedence.LOWEST))
	}

	if p.PeekToken.Type != tokens.RightParenthesis {
		return nil
	}

	p.NextToken()

	return arguments
}

func (p *Parser) parseFunction() ast.Expression {
	token := &ast.Function{
		Token: p.CurrentToken,
	}

	if p.PeekToken.Type != tokens.LeftParenthesis {
		p.AddError(fmt.Sprintf("wrong token. expected=%q. got=%q", "(", p.PeekToken.Literal))
		return nil
	}

	p.NextToken()
	p.NextToken()

	// Parse arguments as identifiers

	var arguments []*ast.Identifier
	for p.CurrentToken.Type == tokens.Identifier || p.CurrentToken.Type == tokens.Comma {
		fmt.Println("?")
		if p.CurrentToken.Type == tokens.Comma {
			continue
		}

		identifier := &ast.Identifier{
			Token: p.CurrentToken,
			Value: p.CurrentToken.Literal,
		}

		arguments = append(arguments, identifier)
		p.NextToken()
	}

	token.Parameters = arguments

	if p.CurrentToken.Type != tokens.RightParenthesis {
		p.AddError(fmt.Sprintf("wrong token. expected=%q. got=%q", ")", p.PeekToken.Literal))
		return nil
	}

	if p.PeekToken.Type != tokens.LeftBrace {
		p.AddError(fmt.Sprintf("wrong token. expected=%q. got=%q", "{", p.PeekToken.Literal))
		return nil
	}

	p.NextToken()

	if p.PeekToken.Type == tokens.EOF {
		p.AddError(fmt.Sprintf("wrong token. expected=%q. got=%q", "}", p.PeekToken.Literal))
		return nil
	}

	p.NextToken()

	for p.CurrentToken.Type != tokens.RightBrace && p.CurrentToken.Type != tokens.EOF {
		stmt := p.parseStatement()

		if stmt != nil {
			token.Body.Statements = append(token.Body.Statements, stmt)
		}

		p.NextToken()
	}

	return token
}

func parseBlock() ast.Expression {
	return &ast.IntegerLiteral{
		Token: tokens.Token{
			Type:    tokens.Integer,
			Literal: "",
			Line:    0,
			Column:  0,
		},
		Value: 0,
	}
}
