package lexer

import (
	"git.kanersps.pw/loop/models/tokens"
	"strconv"
)

type Lexer struct {
	input string
	curPosition int
	peekPosition int
	character byte
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
		l.character = l.input[l.peekPosition]
	}

	l.curPosition = l.peekPosition
	l.peekPosition++
}

func (l *Lexer) PeekCharacter() {

}

func (l *Lexer) NextToken() tokens.Token {
	var token tokens.Token

	l.skipWhitespace()

	switch l.character {
	case '+':
		token = tokens.Token{Type: tokens.Plus, Literal: string(l.character)}
	case ';':
		token = tokens.Token{Type: tokens.Semicolon, Literal: string(l.character)}
	default:
		if isNumber(l.character) {
			value := l.readNumber()
			return tokens.Token{Type: tokens.Integer, Literal: value}
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

func isNumber(input byte) bool {
	_, ok := strconv.Atoi(string(input))

	if ok == nil {
		return true
	} else {
		return false
	}
}
