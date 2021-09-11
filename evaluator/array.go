package evaluator

import (
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
