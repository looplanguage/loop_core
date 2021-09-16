package parser

import "github.com/looplanguage/loop/parser/precedence"

func (p *Parser) peekPrecedence() int {
	if precedence, ok := precedence.Precedences[p.PeekToken.Type]; ok {
		return precedence
	}

	return precedence.LOWEST
}

func (p *Parser) currentPrecedence() int {
	if precedence, ok := precedence.Precedences[p.CurrentToken.Type]; ok {
		return precedence
	}

	return precedence.LOWEST
}
