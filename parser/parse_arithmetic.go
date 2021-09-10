package parser

import (
	"fmt"
	"git.kanersps.pw/loop/models/ast"
	"git.kanersps.pw/loop/models/tokens"
	"git.kanersps.pw/loop/parser/precedence"
	"strconv"
)

func (p *Parser) parseConditionalStatement() ast.Expression {
	token := &ast.ConditionalStatement{
		Token:         p.CurrentToken,
		Condition:     nil,
		ElseCondition: ast.BlockStatement{},
		ElseStatement: nil,
		Body:          ast.BlockStatement{},
	}

	p.NextToken()

	if p.CurrentToken.Type != tokens.LeftParenthesis {
		p.AddError(fmt.Sprintf("wrong token. expected=%q. got=%q", "(", p.CurrentToken.Literal))
		return nil
	}

	token.Condition = p.parseExpression(precedence.LOWEST)

	if p.CurrentToken.Type != tokens.RightParenthesis {
		p.AddError(fmt.Sprintf("wrong token. expected=%q. got=%q", ")", p.CurrentToken.Literal))
		return nil
	}

	p.NextToken()

	if p.CurrentToken.Type != tokens.LeftBrace {
		p.AddError(fmt.Sprintf("wrong token. expected=%q. got=%q", "{", p.CurrentToken.Literal))
		return nil
	}

	p.NextToken()

	token.Body.Statements = p.parseBlockStatement()

	if p.PeekToken.Type == tokens.Else {
		p.NextToken()
		p.NextToken()

		if p.CurrentToken.Type == tokens.If {
			conditionalStatement := p.parseConditionalStatement()
			elseConditional, ok := conditionalStatement.(*ast.ConditionalStatement)
			if ok {
				token.ElseStatement = elseConditional
			} else {
				p.AddError("unable to parse conditional statement, see above.")
				return nil
			}
		} else if p.CurrentToken.Type == tokens.LeftBrace {
			token.ElseCondition.Statements = p.parseBlockStatement()
		}

		if p.CurrentToken.Type != tokens.RightBrace {
			p.AddError(fmt.Sprintf("wrong token. expected=%q. got=%q", "}", p.CurrentToken.Literal))
			return nil
		}

		p.NextToken()
	}

	return token
}

func (p *Parser) parseIntegerLiteral() ast.Expression {

	integer, err := strconv.Atoi(p.CurrentToken.Literal)

	if err != nil {
		fmt.Println(err)
	}

	expr := &ast.IntegerLiteral{
		Token: p.CurrentToken,
		Value: int64(integer),
	}

	if p.peekTokenIs(tokens.Semicolon) {
		p.NextToken()
	}

	return expr
}

func (p *Parser) parseSuffixExpression(left ast.Expression) ast.Expression {
	exp := &ast.SuffixExpression{
		Token:    p.CurrentToken,
		Left:     left,
		Operator: p.CurrentToken.Literal,
	}

	pre := p.currentPrecedence()

	p.NextToken()

	exp.Right = p.parseExpression(pre)

	return exp
}
