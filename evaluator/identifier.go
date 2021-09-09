package evaluator

import (
	"fmt"
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
	return &object.Error{Message: fmt.Sprintf("Identifier %q not defined", identifier.Value)}
}
