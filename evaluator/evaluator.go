package evaluator

import (
	"monkey/ast"
	"monkey/object"
)

var (
	TrueObj  = &object.Boolean{Value: true}
	FalseObj = &object.Boolean{Value: false}
	NullObj  = &object.Null{}
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {

	case *ast.Program:
		return evalStatements(node.Statements)

	case *ast.ExpressionStatement:
		return Eval(node.Expression)

	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}

	case *ast.Boolean:
		if node.Value {
			return TrueObj
		}
		return FalseObj

	}
	return nil
}

func evalStatements(stmt []ast.Statement) object.Object {
	var result object.Object
	for _, statement := range stmt {
		result = Eval(statement)
	}

	return result
}
