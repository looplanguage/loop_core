package evaluator

import (
	"fmt"
	"git.kanersps.pw/loop/models/ast"
	"git.kanersps.pw/loop/models/object"
)

func evalFunction(block ast.BlockStatement, args []*ast.Identifier, env *object.Environment) object.Object {
	obj := &object.Function{
		Parameters:  args,
		Value:       &block,
		Environment: object.CreateEnclosedEnvironment(env),
	}

	return obj
}

func applyFunction(fn object.Object, args []object.Object) object.Object {
	switch fn := fn.(type) {
	case *object.Function:
		env := extendFunctionEnvironment(fn, args)
		evaluated := Eval(fn.Value, env)

		return evaluated
	case *object.BuiltinFunction:
		return fn.Function(args)
	}

	return &object.Error{
		Message: fmt.Sprintf("unknown function %q", fn.Inspect()),
	}
}

func extendFunctionEnvironment(fn *object.Function, args []object.Object) *object.Environment {
	env := object.CreateEnclosedEnvironment(fn.Environment)

	for key, parameter := range fn.Parameters {
		env.Set(parameter.Value, args[key])
	}

	return env
}
