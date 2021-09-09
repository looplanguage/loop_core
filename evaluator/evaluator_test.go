package evaluator

import (
	"git.kanersps.pw/loop/lexer"
	"git.kanersps.pw/loop/models/environment"
	"git.kanersps.pw/loop/models/object"
	"git.kanersps.pw/loop/parser"
	"testing"
)

func TestEval_Integers(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"1", 1},
		{"200", 200},
	}

	test := 1
	for _, tc := range tests {
		evaluated := testEvaluate(tc.input)

		obj := evaluated.(*object.Integer)

		if obj.Type() != object.INTEGER {
			t.Fatalf("(%d/%d) obj.Type is incorrect. expected=%s. got=%s", test, len(tests), object.INTEGER, obj.Type())
		}

		if obj.Value != tc.expected {
			t.Fatalf("(%d/%d) obj.Value is incorrect. expected=%d. got=%d", test, len(tests), tc.expected, obj.Value)
		}

		test++
	}
}

// NOT EVALUATED YET
func TestEval_Variables(t *testing.T) {
	test := "var test = 1"

	_ = testEvaluate(test)

	//if evaluated.Type() != object.
}

func testEvaluate(input string) object.Object {
	l := lexer.Create(input)
	p := parser.Create(l)

	program := p.Parse()

	env := *environment.CreateEnvironment()

	return Eval(program, env)
}
