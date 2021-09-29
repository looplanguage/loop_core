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

type Builtin struct {
	Name    string
	Builtin *BuiltinFunction
}

var Builtins = []Builtin{
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
	{
		"append",
		&BuiltinFunction{
			Function: func(args []Object) Object {
				if len(args) <= 1 {
					return &Error{Message: fmt.Sprintf("wrong number of arguments. expected=%d. got=%d", 2, len(args))}
				}

				if args[0].Type() != "ARRAY" {
					return &Error{Message: fmt.Sprintf("wrong first argument. expected=%q. got=%q", "ARRAY", args[0].Type())}
				}

				array := args[0].(*Array)

				newArray := &Array{Elements: array.Elements}

				args = args[1:]

				newArray.Elements = append(newArray.Elements, args...)

				return newArray
			},
		},
	},
	{
		"slice",
		&BuiltinFunction{
			Function: func(args []Object) Object {
				if len(args) <= 2 {
					return &Error{Message: fmt.Sprintf("wrong number of arguments. expected=%d. got=%d", 3, len(args))}
				}

				array, ok := args[0].(*Array)
				if !ok {
					return &Error{Message: fmt.Sprintf("wrong argument. expected=\"ARRAY\". got=%q", args[0].Type())}
				}

				start, ok := args[1].(*Integer)
				if !ok {
					return &Error{Message: fmt.Sprintf("wrong argument. expected=\"INTEGER\". got=%q", args[0].Type())}
				}

				end, ok := args[2].(*Integer)
				if !ok {
					return &Error{Message: fmt.Sprintf("wrong argument. expected=\"INTEGER\". got=%q", args[0].Type())}
				}

				return &Array{Elements: array.Elements[start.Value:end.Value]}
			},
		},
	},
	{
		"httpServer",
		&BuiltinFunction{Function: func(args []Object) Object {
			return &Null{}
		}},
	},
}
