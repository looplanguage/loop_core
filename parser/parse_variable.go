package parser

import (
	"fmt"
	"git.kanersps.pw/loop/models/ast"
	"git.kanersps.pw/loop/models/tokens"
	"git.kanersps.pw/loop/parser/precedence"
)

func (p *Parser) parseVariable() ast.Statement {
	variable := &ast.VariableDeclaration{Token: p.CurrentToken}

	// Token was (VariableDeclaration) and becomes (Identifier)
	p.NextToken()

	if p.CurrentToken.Type != tokens.Identifier {
		fmt.Println("NIL TOKEN")
		return nil
	}

	name := p.CurrentToken

	// Token was (Identifier) and becomes (Assign)
	p.NextToken()

	if p.CurrentToken.Type != tokens.Assign {
		fmt.Println("NOT ASSIGN")
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
