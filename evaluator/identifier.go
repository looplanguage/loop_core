package evaluator

import (
	"fmt"
	"git.kanersps.pw/loop/models/ast"
	"git.kanersps.pw/loop/models/object"
)

func evalIdentifier(identifier *ast.Identifier, env *object.Environment) object.Object {
	value := env.Get(identifier.Value)

	if value != nil {
		return value
	}

	return &object.Error{Message: fmt.Sprintf("identifier %q not defined", identifier.Value), Line: identifier.Token.Line, Column: identifier.Token.Column}
}
