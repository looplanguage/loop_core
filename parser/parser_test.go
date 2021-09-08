package parser

import (
	"git.kanersps.pw/loop/lexer"
	"testing"
)

func TestParser_NextToken(t *testing.T) {
	input := `1; 1`

	l := lexer.Create(input)
	p := Create(l)

	program := p.Parse()

	if len(program.Statements) != 2 {
		t.Fatalf("len(program.Statements) is not correct. expected=%d. got=%d", 2, len(program.Statements))
	}
}