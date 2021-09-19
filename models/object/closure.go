package object

import "fmt"

type Closure struct {
	Fn   *CompiledFunction
	Free []Object
}

func (c *Closure) Type() string { return CLOSURE }
func (c *Closure) Inspect() string {
	return fmt.Sprintf("Closure[%p]", c)
}
