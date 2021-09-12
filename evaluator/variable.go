package evaluator

import (
	"git.kanersps.pw/loop/models/ast"
	"git.kanersps.pw/loop/models/object"
)

func evalVariableDeclaration(key *ast.Identifier, value ast.Expression, env *object.Environment) object.Object {
	v := Eval(value, env)

	if _, ok := v.(*object.None); ok {
		return &object.Error{
			Message: "unable to assign to NONE",
		}
	}

	return env.Set(key.Value, v)
}
