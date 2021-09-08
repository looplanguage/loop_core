package evaluator

import "git.kanersps.pw/loop/models/object"

func evalSuffixExpression(left object.Object, right object.Object, operator string) object.Object {
	leftValue := left.(*object.Integer)
	rightValue := right.(*object.Integer)

	return &object.Integer{Value: leftValue.Value + rightValue.Value}
}
