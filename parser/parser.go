package parser

import (
	"fmt"
	"github.com/looplanguage/loop/lexer"
	"github.com/looplanguage/loop/models/ast"
	"github.com/looplanguage/loop/models/tokens"
)

type prefixParseFunction func() ast.Expression
type suffixParseFunction func(expression ast.Expression) ast.Expression

type Parser struct {
	lexer *lexer.Lexer

	CurrentToken tokens.Token
	PeekToken    tokens.Token

	prefixParsers map[tokens.TokenType]prefixParseFunction
	suffixParsers map[tokens.TokenType]suffixParseFunction

	Errors []string
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
	parser.registerPrefixParser(tokens.LeftParenthesis, parser.parseGroupedExpression)
	parser.registerPrefixParser(tokens.Func, parser.parseFunction)
	parser.registerPrefixParser(tokens.True, parser.parseBoolean)
	parser.registerPrefixParser(tokens.False, parser.parseBoolean)
	parser.registerPrefixParser(tokens.If, parser.parseConditionalStatement)
	parser.registerPrefixParser(tokens.LeftBracket, parser.parseArray)
	parser.registerPrefixParser(tokens.String, parser.parseString)
	parser.registerPrefixParser(tokens.While, parser.parseWhileLoop)
	parser.registerPrefixParser(tokens.LeftBrace, parser.parseHashmap)
	parser.registerPrefixParser(tokens.Minus, parser.parseMinusInteger)
	parser.registerPrefixParser(tokens.Bang, parser.parseInvertedBoolean)

	// Register suffix parsers
	parser.registerSuffixParser(tokens.Plus, parser.parseSuffixExpression)
	parser.registerSuffixParser(tokens.Asterisk, parser.parseSuffixExpression)
	parser.registerSuffixParser(tokens.Slash, parser.parseSuffixExpression)
	parser.registerSuffixParser(tokens.Minus, parser.parseSuffixExpression)
	parser.registerSuffixParser(tokens.LeftParenthesis, parser.parseCallExpression)
	parser.registerSuffixParser(tokens.Equals, parser.parseSuffixExpression)
	parser.registerSuffixParser(tokens.GreaterThan, parser.parseSuffixExpression)
	parser.registerSuffixParser(tokens.LesserThan, parser.parseSuffixExpression)
	parser.registerSuffixParser(tokens.GreaterEqualsThan, parser.parseSuffixExpression)
	parser.registerSuffixParser(tokens.LesserEqualsThan, parser.parseSuffixExpression)
	parser.registerSuffixParser(tokens.LeftBracket, parser.parseIndexExpression)

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

func (p *Parser) FindError(t tokens.TokenType) {
	//expected := reflect.ValueOf(&book).Elem()

	msg := fmt.Sprintf("Expected %v, got %v instead",
		t, p.PeekToken)
	p.Errors = append(p.Errors, msg)
}

func (p *Parser) expectPeek(t tokens.TokenType) bool {
	if p.peekTokenIs(t) {
		p.NextToken()
		return true
	} else {
		p.FindError(t)
		return false
	}
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

// Add an error
func (p *Parser) AddError(err string) {
	p.Errors = append(p.Errors, fmt.Sprintf("ParserException at %d,%d: %s", p.CurrentToken.Line, p.lexer.CurColumn, err))
}
