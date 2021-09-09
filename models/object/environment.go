package object

type Environment struct {
	variables map[string]Object
	parent    *Environment
}

func CreateEnvironment() *Environment {
	return &Environment{variables: make(map[string]Object)}
}

func CreateEnclosedEnvironment(env *Environment) *Environment {
	return &Environment{
		variables: make(map[string]Object),
		parent:    env,
	}
}

func (e *Environment) Get(key string) Object {
	if e.parent != nil {
		value := e.parent.Get(key)

		if value != nil {
			return value
		}
	}

	if obj, ok := e.variables[key]; ok {
		return obj
	}

	return nil
}

func (e *Environment) Set(key string, value Object) Object {
	e.variables[key] = value

	return e.variables[key]
}
