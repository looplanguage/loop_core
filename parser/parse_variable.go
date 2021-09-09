package parser

import (
	"fmt"
	"git.kanersps.pw/loop/models/ast"
	"git.kanersps.pw/loop/models/tokens"
	"git.kanersps.pw/loop/parser/precedence"
)

func (p *Parser) parseVariable() ast.Statement {
	variable := &ast.VariableDeclaration{Token: p.CurrentToken}

	p.NextToken()

	if p.CurrentToken.Type != tokens.Identifier {
		p.AddError(fmt.Sprintf("wrong token. expected=%q. got=%q", "identifier", p.CurrentToken.Name()))
		return nil
	}

	name := p.CurrentToken

	p.NextToken()

	if p.CurrentToken.Type != tokens.Assign {
		p.AddError(fmt.Sprintf("wrong token. expected=%q. got=%q", "=", p.CurrentToken.Name()))
		return nil
	}

	// Token was (Assign) and becomes (Expression/Value)
	p.NextToken()

	exp := p.parseExpression(precedence.EQUALS)

	variable.Identifier = &ast.Identifier{
		Token: name,
		Value: name.Literal,
	}

	variable.Value = exp

	return variable
}
