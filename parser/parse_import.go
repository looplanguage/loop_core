package parser

import (
	"fmt"
	"github.com/looplanguage/loop/models/ast"
	"github.com/looplanguage/loop/models/tokens"
	"github.com/looplanguage/loop/parser/precedence"
)

func (p *Parser) parseExportStatement() ast.Statement {
	tok := &ast.Export{
		Token: p.CurrentToken,
	}

	if p.PeekToken.Type == tokens.EOF || p.PeekToken.Type == tokens.Semicolon {
		p.AddError(fmt.Sprintf("unknown token. expected=%q. got=%q", "EXPRESSION", p.PeekToken.Name()))
		return nil
	}

	p.NextToken()

	tok.Expression = p.parseExpression(precedence.LOWEST)

	return tok
}

func (p *Parser) parseImportStatement() ast.Statement {
	tok := &ast.Import{
		Token: p.CurrentToken,
	}

	if p.PeekToken.Type != tokens.String {
		p.AddError(fmt.Sprintf("unknown token. expected=%q. got=%q", "STRING", p.PeekToken.Literal))
		return nil
	}

	p.NextToken()

	tok.File = p.CurrentToken.Literal

	if p.PeekToken.Type != tokens.As {
		p.AddError(fmt.Sprintf("unknown token. expected=%q. got=%q", "as", p.PeekToken.Literal))
		return nil
	}

	p.NextToken()

	if p.PeekToken.Type != tokens.Identifier {
		p.AddError(fmt.Sprintf("unknown token. expected=%q. got=%q", "IDENTIFIER", p.PeekToken.Literal))
		return nil
	}

	p.NextToken()

	tok.Identifier = p.CurrentToken.Literal

	return tok
}
