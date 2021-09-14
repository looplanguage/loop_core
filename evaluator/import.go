package evaluator

import (
	"git.kanersps.pw/loop/lexer"
	"git.kanersps.pw/loop/models/ast"
	"git.kanersps.pw/loop/models/object"
	"git.kanersps.pw/loop/parser"
	"io/ioutil"
	"log"
)

func evalExportStatement(exp *ast.Export, env *object.Environment) object.Object {
	return env.Set("__EXPORT", Eval(exp.Expression, env))
}

func evalImportStatement(imp *ast.Import, env *object.Environment) object.Object {
	input, err := ioutil.ReadFile(imp.File)

	if err != nil {
		log.Fatal(err)
	}

	l := lexer.Create(string(input))
	p := parser.Create(l)

	program := p.Parse()

	newEnv := object.CreateEnvironment()
	Eval(program, newEnv)

	if imp.Identifier == "_" {
		var v object.Object
		for key, value := range newEnv.GetAll() {
			v = env.Set(key, value)
		}

		if v == nil {
			v = &object.Null{}
		}

		return v
	} else {
		return env.Set(imp.Identifier, newEnv.Get("__EXPORT"))
	}
}
