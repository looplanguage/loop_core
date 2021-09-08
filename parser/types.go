package parser

import (
	"fmt"
	"git.kanersps.pw/loop/models/ast"
	"git.kanersps.pw/loop/models/tokens"
	"strconv"
)

func (p *Parser) parseIntegerLiteral() ast.Expression {

	integer, err := strconv.Atoi(p.CurrentToken.Literal)

	if err != nil {
		fmt.Println(err)
	}

	expr := &ast.IntegerLiteral{
		Token:   p.CurrentToken,
		Value: int64(integer),
	}

	if p.peekTokenIs(tokens.Semicolon) {
		p.NextToken()
	}

	return expr
}