package environment

import (
	"git.kanersps.pw/loop/models/object"
)

type Environment struct {
	variables map[string]object.Object
}

func CreateEnvironment() *Environment {
	return &Environment{variables: make(map[string]object.Object)}
}

func (e *Environment) Get(key string) object.Object {

	if obj, ok := e.variables[key]; ok {
		return obj
	}

	return nil
}

func (e *Environment) Set(key string, value object.Object) object.Object {
	e.variables[key] = value

	return e.variables[key]
}
