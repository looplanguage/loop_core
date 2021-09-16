package object

import "github.com/looplanguage/loop/models/ast"

type Function struct {
	Parameters  []*ast.Identifier
	Value       *ast.BlockStatement
	Environment *Environment
}

func (f *Function) Type() string    { return FUNCTION }
func (f *Function) Inspect() string { return "function" }
