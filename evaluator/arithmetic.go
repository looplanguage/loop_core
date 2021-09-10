package evaluator

import (
	"fmt"
	"git.kanersps.pw/loop/models/ast"
	"git.kanersps.pw/loop/models/object"
)

func evalSuffixExpression(left object.Object, right object.Object, operator string) object.Object {
	if left, ok := left.(*object.Integer); ok {
		if right, ok := right.(*object.Integer); ok {
			switch operator {
			case "+":
				return &object.Integer{Value: left.Value + right.Value}
			case "*":
				return &object.Integer{Value: left.Value * right.Value}
			case "/":
				return &object.Integer{Value: left.Value / right.Value}
			case "-":
				return &object.Integer{Value: left.Value - right.Value}
			case ">":
				return &object.Boolean{Value: left.Value > right.Value}
			case "<":
				return &object.Boolean{Value: left.Value < right.Value}
			case ">=":
				return &object.Boolean{Value: left.Value >= right.Value}
			case "<=":
				return &object.Boolean{Value: left.Value <= right.Value}
			}
		}
	}

	if operator == "==" {
		if left.Type() == right.Type() && left.Inspect() == right.Inspect() {
			return &TRUE
		}

		return &FALSE
	}

	return &object.Error{Message: fmt.Sprintf("invalid operator. got=%q", operator)}
}

func evalConditionalStatement(condition ast.Expression, ElseCondition ast.BlockStatement, ElseStatement *ast.ConditionalStatement, Body ast.BlockStatement, env *object.Environment) object.Object {
	run := Eval(condition, env)

	if run, ok := run.(*object.Boolean); ok {

		if run.Value {
			var lastEval object.Object

			for _, stmt := range Body.Statements {
				lastEval = Eval(stmt, env)
			}

			return lastEval
		} else {
			if ElseStatement != nil {
				return evalConditionalStatement(ElseStatement.Condition, ElseStatement.ElseCondition, ElseStatement.ElseStatement, ElseStatement.Body, env)
			} else if len(ElseCondition.Statements) > 0 {
				var lastEval object.Object

				for _, stmt := range ElseCondition.Statements {
					lastEval = Eval(stmt, env)
				}

				return lastEval
			}
		}
	}

	return &object.Error{Message: fmt.Sprintf("condition is of invalid type. expected=%q. got=%q", "BOOLEAN", run.Type())}
}
