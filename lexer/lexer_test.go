package lexer

import (
	"git.kanersps.pw/loop/models/tokens"
	"testing"
)

func TestLexer_NextToken(t *testing.T) {
	input := `
100 + 1;
var test = 40;

var fn = fun(a, b) {}

true; false;
`

	expected := []tokens.Token{
		// Integer arithmetic operations
		{Type: tokens.Integer, Literal: "100"},
		{Type: tokens.Plus, Literal: "+"},
		{Type: tokens.Integer, Literal: "1"},
		{Type: tokens.Semicolon, Literal: ";"},

		// Variable declaration
		{Type: tokens.VariableDeclaration, Literal: "var"},
		{Type: tokens.Identifier, Literal: "test"},
		{Type: tokens.Assign, Literal: "="},
		{Type: tokens.Integer, Literal: "40"},
		{Type: tokens.Semicolon, Literal: ";"},

		// Function declaration
		{Type: tokens.VariableDeclaration, Literal: "var"},
		{Type: tokens.Identifier, Literal: "fn"},
		{Type: tokens.Assign, Literal: "="},
		{Type: tokens.Func, Literal: "fun"},
		{Type: tokens.LeftParenthesis, Literal: "("},
		{Type: tokens.Identifier, Literal: "a"},
		{Type: tokens.Comma, Literal: ","},
		{Type: tokens.Identifier, Literal: "b"},
		{Type: tokens.RightParenthesis, Literal: ")"},
		{Type: tokens.LeftBrace, Literal: "{"},
		{Type: tokens.RightBrace, Literal: "}"},

		// Booleans
		{Type: tokens.True, Literal: "true"},
		{Type: tokens.Semicolon, Literal: ";"},
		{Type: tokens.False, Literal: "false"},
		{Type: tokens.Semicolon, Literal: ";"},

		// End of file
		{Type: tokens.EOF, Literal: ""},
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
