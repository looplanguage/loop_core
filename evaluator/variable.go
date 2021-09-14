package evaluator

import (
	"git.kanersps.pw/loop/models/ast"
	"git.kanersps.pw/loop/models/object"
)

func evalVariableDeclaration(key *ast.Identifier, value ast.Expression, env *object.Environment) object.Object {
	v := Eval(value, env)

	set := env.Set(key.Value, v)

	if _, ok := v.(*object.Function); ok {
		return &object.Null{}
	}

	return set
}
