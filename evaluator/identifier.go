package evaluator

import (
	"git.kanersps.pw/loop/models/ast"
	"git.kanersps.pw/loop/models/environment"
	"git.kanersps.pw/loop/models/object"
)

func evalIdentifier(identifier *ast.Identifier, env environment.Environment) object.Object {
	value := env.Get(identifier.Value)

	if value != nil {
		return value
	}

	// TODO: return error
	return &object.Integer{Value: 300}
}
