package evaluator

import (
	"fmt"
	"git.kanersps.pw/loop/models/object"
)

func evalSuffixExpression(left object.Object, right object.Object, operator string) object.Object {
	leftValue := left.(*object.Integer)
	rightValue := right.(*object.Integer)

	switch operator {
	case "+":
		return &object.Integer{Value: leftValue.Value + rightValue.Value}
	case "*":
		return &object.Integer{Value: leftValue.Value * rightValue.Value}
	case "/":
		return &object.Integer{Value: leftValue.Value / rightValue.Value}
	case "-":
		return &object.Integer{Value: leftValue.Value - rightValue.Value}
	case "==":

		if leftValue.Value == rightValue.Value {
			return &TRUE
		}

		return &FALSE
	}

	return &object.Error{Message: fmt.Sprintf("invalid operator. got=%q", operator)}
}
