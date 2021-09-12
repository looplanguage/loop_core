package evaluator

import (
	"git.kanersps.pw/loop/lexer"
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

func TestEval_Variables(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"var test = 1;", 1},
		{"var test = 20 * 2; test;", 40},
		{"var test = 20 * 2; test * 2;", 80},
		{"var test = 20 * 2 - 10; var testtwo = test + 10; testtwo;", 40},
	}

	for c, tc := range tests {
		evaluated := testEvaluate(tc.input)

		if evaluated.Type() != object.INTEGER {
			t.Fatalf("(%d/%d) evaluated.Type is incorrect. expected=%q. got=%q", c+1, len(tests), object.INTEGER, evaluated.Type())
		}

		if evaluated.(*object.Integer).Value != tc.expected {
			t.Fatalf("(%d/%d) evaluated.Value is incorrect. expected=%d. got=%d", c+1, len(tests), tc.expected, evaluated.(*object.Integer).Value)
		}
	}
}

func testEvaluate(input string) object.Object {
	l := lexer.Create(input)
	p := parser.Create(l)

	program := p.Parse()

	env := object.CreateEnvironment()

	return Eval(program, env)
}
