package lexer

import (
	"github.com/looplanguage/loop/models/tokens"
	"strconv"
	"unicode"
)

type Lexer struct {
	input        string
	curPosition  int
	peekPosition int
	character    rune
	line         int
	CurColumn    int
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
		l.CurColumn++
	}

	l.curPosition = l.peekPosition
	l.peekPosition++
}

func (l *Lexer) PeekCharacter() rune {
	if l.peekPosition != 0 {
		return rune(l.input[l.peekPosition])
	}

	return 0
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
	case '{':
		token = createToken(tokens.LeftBrace, l)
	case '}':
		token = createToken(tokens.RightBrace, l)
	case '>':
		token = createToken(tokens.GreaterThan, l)
	case '<':
		token = createToken(tokens.LesserThan, l)
	case ',':
		token = createToken(tokens.Comma, l)
	case '[':
		token = createToken(tokens.LeftBracket, l)
	case ']':
		token = createToken(tokens.RightBracket, l)
	case '"':
		token = createToken(tokens.String, l)
	case ':':
		token = createToken(tokens.Colon, l)
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
			l.CurColumn = 0
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

	if isCharacter(l.character) {
		l.NextCharacter()
	}

	for isCharacter(l.character) || isNumber(l.character) {
		l.NextCharacter()
	}

	return l.input[position:l.curPosition]
}

func (l *Lexer) readString() string {
	position := l.curPosition + 1
	for {
		l.NextCharacter()

		if l.character == '"' || l.character == 0 {
			break
		}
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
	if unicode.IsLetter(input) || input == '_' {
		return true
	}

	return false
}

func createToken(token tokens.TokenType, l *Lexer) tokens.Token {
	literal := string(l.character)
	column := l.CurColumn

	if isNumber(l.character) {
		literal = l.readNumber()
	} else if isCharacter(l.character) {
		literal = l.readIdentifier()
		token = tokens.FindKeyword(literal)
	} else {
		switch l.character {
		case '=':
			if l.PeekCharacter() == '=' {
				token = tokens.Equals
				literal = "=="
				l.NextCharacter()
			}
		case '>':
			if l.PeekCharacter() == '=' {
				token = tokens.Equals
				literal = ">="
				l.NextCharacter()
			}
		case '<':
			if l.PeekCharacter() == '=' {
				token = tokens.Equals
				literal = "<="
				l.NextCharacter()
			}
		case '"':
			literal = l.readString()
		}
	}

	return tokens.Token{Type: token, Literal: literal, Line: l.line, Column: column - 1}
}
