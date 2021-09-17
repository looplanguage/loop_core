package parser

import (
	"fmt"
	"github.com/looplanguage/loop/models/ast"
	"github.com/looplanguage/loop/models/tokens"
	"github.com/looplanguage/loop/parser/precedence"
	"strconv"
)

func (p *Parser) parseConditionalStatement() ast.Expression {
	token := &ast.ConditionalStatement{
		Token: p.CurrentToken,
	}

	p.NextToken()

	if p.CurrentToken.Type != tokens.LeftParenthesis {
		p.AddError(fmt.Sprintf("wrong token. expected=%q. got=%q", "(", p.CurrentToken.Literal))
		return nil
	}

	p.NextToken()

	token.Condition = p.parseExpression(precedence.LOWEST)

	p.NextToken()

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

	token.Body = p.parseBlockStatement()

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
			p.NextToken()
			token.ElseCondition = p.parseBlockStatement()
		}
	}

	return token
}

func (p *Parser) parseIntegerLiteral() ast.Expression {

	integer, err := strconv.ParseInt(p.CurrentToken.Literal, 0, 64)

	if err != nil {
		fmt.Println(err)
	}

	expr := &ast.IntegerLiteral{
		Token: p.CurrentToken,
		Value: integer,
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
