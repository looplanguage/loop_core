package parser

import (
	"git.kanersps.pw/loop/lexer"
	"git.kanersps.pw/loop/models/ast"
	"git.kanersps.pw/loop/models/tokens"
)

type prefixParseFunction func() ast.Expression
type suffixParseFunction func(expression ast.Expression) ast.Expression

type Parser struct {
	lexer *lexer.Lexer

	CurrentToken tokens.Token
	PeekToken    tokens.Token

	prefixParsers map[tokens.TokenType]prefixParseFunction
	suffixParsers map[tokens.TokenType]suffixParseFunction
}

func Create(l *lexer.Lexer) *Parser {
	parser := &Parser{
		lexer:         l,
		prefixParsers: make(map[tokens.TokenType]prefixParseFunction),
		suffixParsers: make(map[tokens.TokenType]suffixParseFunction),
	}

	// Register prefix parsers
	parser.registerPrefixParser(tokens.Integer, parser.parseIntegerLiteral)
	parser.registerPrefixParser(tokens.Identifier, parser.parseIdentifier)

	// Register suffix parsers
	parser.registerSuffixParser(tokens.Plus, parser.parseSuffixExpression)
	parser.registerSuffixParser(tokens.Asterisk, parser.parseSuffixExpression)
	parser.registerSuffixParser(tokens.Slash, parser.parseSuffixExpression)
	parser.registerSuffixParser(tokens.Minus, parser.parseSuffixExpression)

	// Call twice to fill CurrentToken & PeekToken
	parser.NextToken()
	parser.NextToken()

	return parser
}

// Actually parse the program into an Abstract Syntax Tree
func (p *Parser) Parse() *ast.Program {
	program := &ast.Program{Statements: []ast.Statement{}}

	// Go through all tokens -> convert them to statements -> add to program.Statements
	for p.CurrentToken.Type != tokens.EOF {
		stmt := p.parseStatement()

		program.Statements = append(program.Statements, stmt)

		p.NextToken()
	}

	return program
}

// Register different parsers (prefix & suffix)
func (p *Parser) registerPrefixParser(tokenType tokens.TokenType, function prefixParseFunction) {
	p.prefixParsers[tokenType] = function
}

func (p *Parser) registerSuffixParser(tokenType tokens.TokenType, function suffixParseFunction) {
	p.suffixParsers[tokenType] = function
}

// Check next token
func (p *Parser) peekTokenIs(tokenType tokens.TokenType) bool {
	if p.PeekToken.Type == tokenType {
		return true
	}

	return false
}

// Peek precedence

func (p *Parser) NextToken() {
	p.CurrentToken = p.PeekToken
	p.PeekToken = p.lexer.NextToken()
}
