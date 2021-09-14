package object

import (
	"git.kanersps.pw/loop/models/ast"
)

type Function struct {
	Parameters  []*ast.Identifier
	Value       *ast.BlockStatement
	Environment *Environment
}

func (f *Function) Type() string    { return FUNCTION }
func (f *Function) Inspect() string { return "function" }
