package lexer

import (
	"fmt"
	"git.kanersps.pw/loop/models/tokens"
	"strconv"
	"unicode"
)

type Lexer struct {
	input        string
	curPosition  int
	peekPosition int
	character    rune
	line         int
	curColumn    int
}

func Create(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.NextCharacter()

	return lexer
}

func (l *Lexer) NextCharacter() {
	if l.peekPosition >= len(l.input) {
		l.character = 0
	} else {
		l.character = rune(l.input[l.peekPosition])
	}

	l.curPosition = l.peekPosition
	l.curColumn++
	l.peekPosition++
}

func (l *Lexer) PeekCharacter() {

}

func (l *Lexer) NextToken() tokens.Token {
	var token tokens.Token

	l.skipWhitespace()

	switch l.character {
	case '+':
		token = createToken(tokens.Plus, l)
	case '*':
		token = createToken(tokens.Asterisk, l)
	case '/':
		token = createToken(tokens.Slash, l)
	case '-':
		token = createToken(tokens.Minus, l)
	case ';':
		token = createToken(tokens.Semicolon, l)
	case '=':
		token = createToken(tokens.Assign, l)
	case '(':
		token = createToken(tokens.LeftParenthesis, l)
	case ')':
		token = createToken(tokens.RightParenthesis, l)
	default:
		if isNumber(l.character) {
			return createToken(tokens.Integer, l)
		} else if isCharacter(l.character) {
			return createToken(tokens.Identifier, l)
		} else {
			token = tokens.Token{
				Type:    tokens.EOF,
				Literal: "",
			}
		}
	}

	l.NextCharacter()

	return token
}

func (l *Lexer) skipWhitespace() {
	for l.character == ' ' || l.character == '\t' || l.character == '\n' || l.character == '\r' {
		if l.character == '\n' || l.character == '\r' {
			l.line++
			l.curColumn = 0
		}

		l.NextCharacter()
	}
}

func (l *Lexer) readNumber() string {
	position := l.curPosition

	for isNumber(l.character) {
		l.NextCharacter()
	}

	return l.input[position:l.curPosition]
}

func (l *Lexer) readIdentifier() string {
	position := l.curPosition

	for isCharacter(l.character) {
		l.NextCharacter()
	}

	return l.input[position:l.curPosition]
}

func isNumber(input rune) bool {
	_, ok := strconv.Atoi(string(input))

	if ok == nil {
		return true
	} else {
		return false
	}
}

func isCharacter(input rune) bool {
	if unicode.IsLetter(input) {
		return true
	}

	return false
}

func createToken(token tokens.TokenType, l *Lexer) tokens.Token {
	literal := string(l.character)

	if isNumber(l.character) {
		literal = l.readNumber()
	} else if isCharacter(l.character) {
		literal = l.readIdentifier()
		token = tokens.FindKeyword(literal)
	}

	fmt.Println(l.curColumn)
	return tokens.Token{Type: token, Literal: literal, Line: l.line, Column: l.curColumn}
}
