package parser

import (
	"git.kanersps.pw/loop/lexer"
	"testing"
)

func TestParser_NextToken(t *testing.T) {
	input := `
1;
var test = 1;
test;
true; false;
`

	l := lexer.Create(input)
	p := Create(l)

	program := p.Parse()

	if len(program.Statements) != 5 {
		t.Fatalf("len(program.Statements) is not correct. expected=%d. got=%d", 2, len(program.Statements))
	}
}
