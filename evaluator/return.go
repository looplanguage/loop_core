package evaluator

import (
	"git.kanersps.pw/loop/models/ast"
	"git.kanersps.pw/loop/models/object"
)

func evalReturn(value ast.Expression) object.Object {
	r := &object.Return{Value: value}

	return r
}
