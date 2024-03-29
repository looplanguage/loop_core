package object

type Environment struct {
	variables map[string]Object
	parent    *Environment
	FileRoot  string
}

func CreateEnvironment(fileroot string) *Environment {
	return &Environment{variables: make(map[string]Object), FileRoot: fileroot}
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
	if e.parent != nil {
		if _, ok := e.parent.variables[key]; ok {
			return e.parent.Set(key, value)
		}
	}

	e.variables[key] = value

	return e.variables[key]
}

func (e *Environment) GetAll() map[string]Object {
	return e.variables
}
