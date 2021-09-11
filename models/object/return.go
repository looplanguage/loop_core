package object

import "git.kanersps.pw/loop/models/ast"

type Return struct {
	Value ast.Expression
}

func (r *Return) Type() string    { return RETURN }
func (r *Return) Inspect() string { return r.Value.String() }
