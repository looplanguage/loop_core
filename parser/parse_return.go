package parser

import (
	"github.com/looplanguage/loop/models/ast"
	"github.com/looplanguage/loop/parser/precedence"
)

func (p *Parser) parseReturn() ast.Statement {
	ret := &ast.Return{Token: p.CurrentToken}

	p.NextToken()

	exp := p.parseExpression(precedence.LOWEST)

	if exp == nil {
		p.AddError("expression was nil, see above.")
		return nil
	}

	ret.Value = exp

	return ret
}
