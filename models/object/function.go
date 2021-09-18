package object

import (
	"fmt"
	"github.com/looplanguage/loop/models/ast"
)

type Function struct {
	Parameters  []*ast.Identifier
	Value       *ast.BlockStatement
	Environment *Environment
}

func (f *Function) Type() string    { return FUNCTION }
func (f *Function) Inspect() string { return "function" }

type CompiledFunction struct {
	Instructions []byte
}

func (cf *CompiledFunction) Type() string    { return COMPILED_FUNCTION }
func (cf *CompiledFunction) Inspect() string { return fmt.Sprintf("CompiledFunction[%p]", cf) }
