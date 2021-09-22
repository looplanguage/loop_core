package parser

import (
	"fmt"
	"github.com/looplanguage/loop/lexer"
	"log"
	"testing"
)

func TestParser_NextToken(t *testing.T) {
	input := `
1;
var test = 1;
test;
true; false;

if(true) {} else if(true) {}else if(true) {} else {};
1000;
i = 100
`

	l := lexer.Create(input)
	p := Create(l)

	program := p.Parse()

	if len(p.Errors) != 0 {
		for _, e := range p.Errors {
			fmt.Println(e)
		}
		log.Fatalln("Errors! ^")
	}

	if len(program.Statements) != 8 {
		t.Fatalf("len(program.Statements) is not correct. expected=%d. got=%d", 7, len(program.Statements))
	}
}
