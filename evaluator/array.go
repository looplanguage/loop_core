package evaluator

import (
	"fmt"
	"git.kanersps.pw/loop/models/ast"
	"git.kanersps.pw/loop/models/object"
)

func evalArray(arr *ast.Array, env *object.Environment) object.Object {
	var values []object.Object
	array := &object.Array{}

	for _, exp := range arr.Elements {
		val := Eval(exp, env)

		if val != nil {
			values = append(values, val)
		}
	}

	array.Elements = values

	return array
}

func evalIndexExpression(value ast.Expression, index ast.Expression, env *object.Environment) object.Object {
	val := Eval(value, env)
	idx := Eval(index, env)

	if array, ok := val.(*object.Array); ok {
		if _index, ok := idx.(*object.Integer); ok {
			return array.Elements[_index.Value]
		} else {
			return &object.Error{
				Message: fmt.Sprintf("unknown value to index. got=%q", value.TokenLiteral()),
			}
		}
	} else {
		return &object.Error{
			Message: fmt.Sprintf("unknown value to index. got=%q", val.Inspect()),
		}
	}
}
