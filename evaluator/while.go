package evaluator

import (
	"fmt"
	"git.kanersps.pw/loop/models/ast"
	"git.kanersps.pw/loop/models/object"
)

func evalWhile(condition ast.Expression, block *ast.BlockStatement, env *object.Environment) object.Object {
	exp := Eval(condition, env)

	bool, ok := exp.(*object.Boolean)

	if !ok {
		return &object.Error{Message: fmt.Sprintf("wrong expression value. expected=\"BOOLEAN\". got=%q", exp.Type())}
	}

	var rValue object.Object

	for bool.Value {
		bool = Eval(condition, env).(*object.Boolean)

		rValue = Eval(block, env)
	}

	rValue = &object.None{}

	return rValue
}
