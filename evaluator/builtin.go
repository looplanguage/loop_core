package evaluator

import (
	"fmt"
	"git.kanersps.pw/loop/models/object"
)

var BuiltinFunctions = map[string]*object.BuiltinFunction{
	"len": {
		Function: func(args []object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: fmt.Sprintf("not enough arguments to builtin function \"len\". expected=1. got=%d", len(args))}
			}

			switch arg := args[0].(type) {
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			}

			return &object.Error{Message: fmt.Sprintf("wrong type as argument for \"len\". got=%T", args[0])}
		},
	},
}
