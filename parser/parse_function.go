package parser

import (
	"fmt"
	"git.kanersps.pw/loop/models/ast"
	"git.kanersps.pw/loop/models/tokens"
	"git.kanersps.pw/loop/parser/precedence"
)

func (p *Parser) parseCallExpression(fn ast.Expression) ast.Expression {
	params := p.parseCallArguments()

	return &ast.CallExpression{
		Token:      p.CurrentToken,
		Function:   fn,
		Parameters: params,
	}
}

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
	for p.CurrentToken.Type == tokens.Identifier {
		identifier := &ast.Identifier{
			Token: p.CurrentToken,
			Value: p.CurrentToken.Literal,
		}

		arguments = append(arguments, identifier)
		p.NextToken()

		if p.CurrentToken.Type != tokens.RightParenthesis {
			p.NextToken()
		}
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

	token.Body.Statements = p.parseBlockStatement()

	return token
}

func (p *Parser) parseBlockStatement() []ast.Statement {
	var statements []ast.Statement

	for p.CurrentToken.Type != tokens.RightBrace && p.CurrentToken.Type != tokens.EOF {
		stmt := p.parseStatement()

		if stmt != nil {
			statements = append(statements, stmt)
		}

		p.NextToken()
	}

	return statements
}
