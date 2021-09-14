package object

import "fmt"

type Export struct {
	Export Object
}

func (e *Export) Type() string    { return EXPORT }
func (e *Export) Inspect() string { return fmt.Sprintf("export %s", e.Export.Inspect()) }
