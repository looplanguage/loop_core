package object

type BuiltinFunction struct {
	Function func(args []Object) Object
}

func (bf *BuiltinFunction) Type() string    { return BUILTIN }
func (bf *BuiltinFunction) Inspect() string { return "builtin fun" }
