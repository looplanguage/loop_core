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

// TODO: Clean this mess up
func evalIndexExpression(value ast.Expression, index ast.Expression, env *object.Environment) object.Object {
	val := Eval(value, env)
	idx := Eval(index, env)

	if val.Type() == object.ARRAY {
		if array, ok := val.(*object.Array); ok {
			if _index, ok := idx.(*object.Integer); ok {
				// TODO: Fix null reference possibility
				return array.Elements[_index.Value]
			} else {
				return &object.Error{
					Message: fmt.Sprintf("unknown value to index. got=%q", value.String()),
				}
			}
		} else {
			return &object.Error{
				Message: fmt.Sprintf("unknown value to index. got=%q", val.Inspect()),
			}
		}
	} else if val.Type() == object.HASHMAP {
		if hashmap, ok := val.(*object.HashMap); ok {
			if hashkey, ok := idx.(object.Hash); ok {
				val, ok := hashmap.Pairs[hashkey.Hash()]

				if ok {
					return val.Value
				} else {
					return &object.Error{
						Message: fmt.Sprintf("value with index does not exist. got=%q", idx.Inspect()),
					}
				}
			} else {
				return &object.Error{
					Message: fmt.Sprintf("unknown index. got=%q", hashkey),
				}
			}
		} else {
			return &object.Error{
				Message: fmt.Sprintf("unknown value to index. got=%q", val.Inspect()),
			}
		}
	}

	return &object.Error{Message: fmt.Sprintf("unknown value to index. got=%q", val.Type())}
}
