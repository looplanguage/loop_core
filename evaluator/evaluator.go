package evaluator

import (
	"git.kanersps.pw/loop/models/ast"
	"git.kanersps.pw/loop/models/environment"
	"git.kanersps.pw/loop/models/object"
)

func Eval(node ast.Node, env environment.Environment) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node.Statements, env)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)
	case *ast.VariableDeclaration:
		return evalVariableDeclaration(node.Identifier, node.Value, env)
	case *ast.Identifier:
		return evalIdentifier(node, env)
	case *ast.SuffixExpression:
		left := Eval(node.Left, env)
		right := Eval(node.Right, env)

		return evalSuffixExpression(left, right, "")
	}

	return nil
}

func evalProgram(statements []ast.Statement, env environment.Environment) object.Object {
	var result object.Object

	for _, statement := range statements {
		result = Eval(statement, env)
	}

	return result
}
