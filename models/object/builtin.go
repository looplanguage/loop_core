package object

import "fmt"

type BuiltinFunction struct {
	Function func(args []Object) Object
}

func (bf *BuiltinFunction) Type() string    { return BUILTIN }
func (bf *BuiltinFunction) Inspect() string { return "builtin fun" }

func GetBuiltinByName(name string) *BuiltinFunction {
	for _, def := range Builtins {
		if def.Name == name {
			return def.Builtin
		}
	}

	return nil
}

var Builtins = []struct {
	Name    string
	Builtin *BuiltinFunction
}{
	{
		"len",
		&BuiltinFunction{
			Function: func(args []Object) Object {
				if len(args) != 1 {
					return &Error{Message: fmt.Sprintf("wrong number of arguments. expected=%d. got=%d", 1, len(args))}
				}

				switch arg := args[0].(type) {
				case *Array:
					return &Integer{Value: int64(len(arg.Elements))}
				case *String:
					return &Integer{Value: int64(len(arg.Value))}
				}

				return &Error{Message: fmt.Sprintf("incorrect argument type, can not iterate. got=%q", args[0].Type())}
			},
		},
	},
	{
		"print",
		&BuiltinFunction{
			Function: func(args []Object) Object {
				for _, arg := range args {
					fmt.Println(arg.Inspect())
				}

				return nil
			},
		},
	},
}
