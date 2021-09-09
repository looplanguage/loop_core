package evaluator

import (
	"git.kanersps.pw/loop/models/ast"
	"git.kanersps.pw/loop/models/object"
)

func evalVariableDeclaration(key *ast.Identifier, value ast.Expression, env *object.Environment) object.Object {
	return env.Set(key.Value, Eval(value, env))
}
