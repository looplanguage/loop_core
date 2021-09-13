package parser

import (
	"fmt"
	"git.kanersps.pw/loop/models/ast"
	"git.kanersps.pw/loop/models/tokens"
	"git.kanersps.pw/loop/parser/precedence"
)

func (p *Parser) parseWhileLoop() ast.Expression {
	while := &ast.While{
		Token: p.CurrentToken,
	}

	if p.PeekToken.Type != tokens.LeftParenthesis {
		p.AddError(fmt.Sprintf("wrong token. expected=%q. got=%q", "(", p.PeekToken.Literal))
		return nil
	}

	p.NextToken()
	p.NextToken()

	while.Condition = p.parseExpression(precedence.LOWEST)

	p.NextToken()

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
	fmt.Println(p.CurrentToken.Literal)

	while.Block = &ast.BlockStatement{}

	while.Block.Statements = p.parseBlockStatement()

	return while
}
