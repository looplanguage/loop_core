package lexer

import (
	"git.kanersps.pw/loop/models/tokens"
	"testing"
)

func TestLexer_NextToken(t *testing.T) {
	input := `100 + 1`

	expected := []tokens.Token {
		{ Type: tokens.Integer, Literal: "100"},
		{ Type: tokens.Plus, Literal: "+"},
		{ Type: tokens.Integer, Literal: "1"},
		{ Type: tokens.EOF, Literal: ""},
	}

	l := Create(input)

	token := l.NextToken()

	current := 1
	for _, e := range expected {
		if e.Type != token.Type {
			t.Fatalf("(%d/%d) lexer failed parsing correct token (type). expected=%q. got=%q", current, len(expected), e.Name(), token.Name())
		}

		if e.Literal != token.Literal {
			t.Fatalf("(%d/%d) lexer failed parsing correct token (value). expected=%q. got=%q", current, len(expected), e.Literal, token.Literal)
		}

		current++
		token = l.NextToken()
	}
}
