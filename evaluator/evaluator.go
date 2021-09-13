package evaluator

import (
	"git.kanersps.pw/loop/models/ast"
	"git.kanersps.pw/loop/models/object"
)

var (
	TRUE  = object.Boolean{Value: true}
	FALSE = object.Boolean{Value: false}
)

func Eval(node ast.Node, env *object.Environment) object.Object {
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
	case *ast.Function:
		return evalFunction(node.Body, node.Parameters, env)
	case *ast.CallExpression:
		function := Eval(node.Function, env)
		if function.Type() == object.ERROR {
			return function
		}

		args := evalExpressions(node.Parameters, env)

		if len(args) == 1 && args[0].Type() == object.ERROR {
			return args[0]
		}

		return applyFunction(function, args)
	case *ast.SuffixExpression:
		left := Eval(node.Left, env)
		right := Eval(node.Right, env)

		return evalSuffixExpression(left, right, node.Operator)
	case *ast.BlockStatement:
		var val object.Object

		for _, exp := range node.Statements {
			val = Eval(exp, env)

			if ret, earlyReturn := val.(*object.Return); earlyReturn {
				value := Eval(ret.Value, env)

				return value
			}
		}

		if val == nil {
			val = &object.None{}
		}

		return val
	case *ast.ConditionalStatement:
		return evalConditionalStatement(node.Condition, node.ElseCondition, node.ElseStatement, node.Body, env)
	case *ast.Boolean:
		return boolToObject(node.Value)
	case *ast.Array:
		return evalArray(node, env)
	case *ast.IndexExpression:
		return evalIndexExpression(node.Value, node.Index, env)
	case *ast.Return:
		return evalReturn(node.Value)
	case *ast.String:
		return evalString(node.Value)
	case *ast.While:
		return evalWhile(node.Condition, node.Block, env)
	}

	return nil
}

func evalExpressions(parameters []ast.Expression, env *object.Environment) []object.Object {
	var params []object.Object

	for _, param := range parameters {
		exp := Eval(param, env)

		if exp != nil {
			params = append(params, exp)
		}
	}

	return params
}

func evalProgram(statements []ast.Statement, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range statements {
		result = Eval(statement, env)
	}

	return result
}

func boolToObject(b bool) *object.Boolean {
	if b {
		return &TRUE
	}

	return &FALSE
}
