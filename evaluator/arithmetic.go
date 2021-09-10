package evaluator

import (
	"fmt"
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
