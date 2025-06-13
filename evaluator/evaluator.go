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

	case *ast.PrefixExpression:
		right := Eval(node.Right)
		return evalPrefixExpression(node.Operator, right)

	case *ast.InfixExpression:
		left := Eval(node.Left)
		right := Eval(node.Right)
		return evalInfixExpression(node.Operator, left, right)
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

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusPrefixOperatorExpression(right)
	case "==":

	}

	return NullObj
}

func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case TrueObj:
		return FalseObj
	case FalseObj:
		return TrueObj
	case NullObj:
		return TrueObj
	default:
		return FalseObj
	}
}

func evalMinusPrefixOperatorExpression(right object.Object) object.Object {

	value, ok := right.(*object.Integer)
	if !ok {
		return NullObj
	}

	return &object.Integer{Value: -value.Value}
}

func evalInfixExpression(operator string, left, right object.Object) object.Object {
	switch {
	case left.Type() == object.IntegerTypeObj && right.Type() == object.IntegerTypeObj:
		return evalIntegerInfixExpression(operator, left, right)
	case operator == "==":
		return nativeBoolToBooleanObj(left == right)
	case operator == "!=":
		return nativeBoolToBooleanObj(left != right)
	default:
		return NullObj
	}
}

func evalIntegerInfixExpression(operator string, left, right object.Object) object.Object {
	leftValue := left.(*object.Integer).Value
	rightValue := right.(*object.Integer).Value

	switch operator {
	case "+":
		return &object.Integer{Value: leftValue + rightValue}
	case "-":
		return &object.Integer{Value: leftValue - rightValue}
	case "*":
		return &object.Integer{Value: leftValue * rightValue}
	case "/":
		return &object.Integer{Value: leftValue / rightValue}
	case "<":
		return nativeBoolToBooleanObj(leftValue < rightValue)
	case ">":
		return nativeBoolToBooleanObj(leftValue > rightValue)
	case "==":
		return nativeBoolToBooleanObj(leftValue == rightValue)
	case "!=":
		return nativeBoolToBooleanObj(leftValue != rightValue)

	}

	return NullObj
}

func nativeBoolToBooleanObj(input bool) *object.Boolean {
	if input {
		return TrueObj
	}
	return FalseObj
}
