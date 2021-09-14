package parser

import (
	"fmt"
	"git.kanersps.pw/loop/models/ast"
	"git.kanersps.pw/loop/models/tokens"
	"git.kanersps.pw/loop/parser/precedence"
)

func (p *Parser) parseHashmap() ast.Expression {
	hash := &ast.Hashmap{Token: p.CurrentToken}
	hash.Values = make(map[ast.Expression]ast.Expression)

	for !p.peekTokenIs(tokens.RightBrace) {
		p.NextToken()
		key := p.parseExpression(precedence.LOWEST)

		if p.PeekToken.Type != tokens.Colon {
			fmt.Println(fmt.Sprintf("wrong token. expected=%q. got=%q", ":", p.PeekToken.Literal))
			return nil
		}

		p.NextToken()
		p.NextToken()

		value := p.parseExpression(precedence.LOWEST)
		hash.Values[key] = value

		if !p.peekTokenIs(tokens.RightBrace) && !p.peekTokenIs(tokens.Comma) {
			fmt.Println(fmt.Sprintf("wrong token. expected=%q. got=%q", "}", p.PeekToken.Literal))
			return nil
		}

		if p.peekTokenIs(tokens.Comma) {
			p.NextToken()
		}
	}

	if !p.peekTokenIs(tokens.RightBrace) {
		fmt.Println(fmt.Sprintf("wrong token. expected=%q. got=%q", "}", p.PeekToken.Literal))
	}

	p.NextToken()

	return hash
}
