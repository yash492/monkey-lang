package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/lexer"
	"testing"
)

func TestLetStatments(t *testing.T) {
	input := `
		let x = 5;
		let y = 10;
		let foobar = 77777;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statments does not contain 3 statements. got %d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		testLetStatement(t, stmt, tt.expectedIdentifier)
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral() not 'let' got %s", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got%T", s)
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.TokenLiteral() not %s, got=%s", name, letStmt.Name.TokenLiteral())
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not %s, got=%s", name, letStmt.Name.Value)
		return false
	}

	return true
}

func TestReturnStatments(t *testing.T) {
	input := `
		return xyz;
		return true;
		return 5;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statments does not contain 3 statements. got %d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("s not *ast.ReturnStatement. got%T", stmt)
		}

		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmy.TokenLiteral() not %s, got=%s", "return", returnStmt.TokenLiteral())
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program does not have enough statements, got %d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statement[0] is not ast.ExpressionStatement, got=%T", program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("exp not ast.Identifier, got %T", stmt.Expression)
	}

	if ident.Value != "foobar" {
		t.Fatalf("ident.Value not %s, got=%s", "foobar", ident.Value)
	}

	if ident.TokenLiteral() != "foobar" {
		t.Fatalf("ident.TokenLiteral not %s, got=%s", "foobar", ident.TokenLiteral())
	}

}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errors))

	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "5;"
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program does not contain %d statements, got=%d/n", 1, len(program.Statements))
	}

	expr, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("ast.Expression expected, got %T", program.Statements[0])
	}

	integerExpr, ok := expr.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("ast.IntegerLiteral expected, got %T", expr)
	}

	if integerExpr.Value != 5 {
		t.Fatalf("integerExpr.Value not %d, got = %d", 5, integerExpr.Value)
	}

	if integerExpr.TokenLiteral() != "5" {
		t.Fatalf("integerExpr.TokenLiteral() not %s, got = %s", "5;", integerExpr.TokenLiteral())
	}
}

func TestParsingPrefixExpressions(t *testing.T) {
	prefixTests := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{
			"!5;", "!", 5,
		},
		{
			"!15;", "!", 15,
		},
	}

	for _, tt := range prefixTests {
		l := lexer.New(tt.input)
		p := New(l)
		// This parses the program and converts the code into AST
		program := p.ParseProgram()
		checkParserErrors(t, p)
		if len(program.Statements) != 1 {
			t.Fatalf("program does not contain %d statements, got=%d/n", 1, len(program.Statements))
		}

		expr, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("ast.Expression expected, got %T", program.Statements[0])
		}

		exp, ok := expr.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("ast.PrefixExpression expected, got %T", expr)
		}

		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not %s, got %s", tt.operator, exp.Operator)
		}

		if !testLiteralExpression(t, exp.Right, tt.integerValue) {
			return
		}
	}
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	integerExpr, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("ast.IntegerLiteral expected, got %T", integerExpr)
		return false
	}

	if integerExpr.Value != value {
		t.Fatalf("integerExpr.Value not %d, got = %d", integerExpr.Value, value)
		return false
	}

	return true
}

func TestParsingInfixExpressions(t *testing.T) {
	prefixTests := []struct {
		input      string
		leftValue  any
		operator   string
		rightValue any
	}{
		{
			"5 + 5;", 5, "+", 5,
		},
		{
			"5 - 5;", 5, "-", 5,
		},
		{
			"5 * 5;", 5, "*", 5,
		},
		{
			"5 / 5;", 5, "/", 5,
		},
		{
			"5 > 5;", 5, ">", 5,
		},
		{
			"5 < 5;", 5, "<", 5,
		},
		{
			"5 == 5;", 5, "==", 5,
		},
		{
			"5 != 5;", 5, "!=", 5,
		},
		{
			"true == true", true, "==", true,
		},
	}

	for _, tt := range prefixTests {
		l := lexer.New(tt.input)
		p := New(l)
		// This parses the program and converts the code into AST
		program := p.ParseProgram()
		checkParserErrors(t, p)
		if len(program.Statements) != 1 {
			t.Fatalf("program does not contain %d statements, got=%d/n", 1, len(program.Statements))
		}

		expr, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("ast.Expression expected, got %T", program.Statements[0])
		}

		exp, ok := expr.Expression.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("ast.InfixExpression expected, got %T", expr)
		}

		if !testLiteralExpression(t, exp.Left, tt.leftValue) {
			return
		}

		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not %s, got %s", tt.operator, exp.Operator)
		}

		if !testLiteralExpression(t, exp.Right, tt.rightValue) {
			return
		}
	}
}

func TestOperatorPrecedenceParsing(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{
			input:  "a * (b + c) -d",
			output: "((a * (b + c)) - d)",
		},
		{
			input:  "a + b * c - d / e - f",
			output: "(((a + (b * c)) - (d / e)) - f)",
		},
		{
			input:  "3 + 4; -5 * 5",
			output: "(3 + 4)((-5) * 5)",
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)

		actual := p.ParseProgram()
		checkParserErrors(t, p)

		if tt.output != actual.String() {
			t.Errorf("expected=%q; got=%q", tt.output, actual.String())
		}
	}
}

func testIdentifier(t *testing.T, exp ast.Expression, value string) bool {
	ident, ok := exp.(*ast.Identifier)
	if !ok {
		t.Errorf("exp not *ast.Identifier, got=%T", exp)
		return false
	}

	if ident.Value != value {
		t.Errorf("ident.Value not %s. got=%s", value, ident.Value)
		return false
	}

	if ident.TokenLiteral() != value {
		t.Errorf("ident.TokenLiteral not %s. got=%s", value, ident.TokenLiteral())
		return false
	}

	return true
}

func testLiteralExpression(t *testing.T, exp ast.Expression, expected any) bool {
	switch v := expected.(type) {
	case int:
		return testIntegerLiteral(t, exp, int64(v))

	case bool:
		return testBooleanLiteral(t, exp, v)

	case int64:
		return testIntegerLiteral(t, exp, v)

	case string:
		return testIdentifier(t, exp, v)
	}

	t.Errorf("type of exp not handled. got=%s", exp)

	return false

}

func testInfixExpression(
	t *testing.T,
	exp ast.Expression,
	left any,
	operator string,
	right any,
) bool {

	opExp, ok := exp.(*ast.InfixExpression)
	if !ok {
		t.Errorf("exp is not ast.InfixExpression, got=%T(%s)", exp, exp)
	}

	if !testLiteralExpression(t, opExp.Left, left) {
		return false
	}

	if opExp.Operator != operator {
		t.Errorf("exp.Operator is not '%s. got=%q", operator, opExp.Operator)
		return false
	}

	if !testLiteralExpression(t, opExp.Right, right) {
		return false
	}

	return true
}

func testBooleanLiteral(t *testing.T, exp ast.Expression, value bool) bool {
	bo, ok := exp.(*ast.Boolean)
	if !ok {
		t.Errorf("exp not *ast.Boolean. got=%T", exp)
		return false
	}

	if bo.Value != value {
		t.Errorf("bo.Value not %t. got=%t", value, bo.Value)
		return false
	}

	if bo.TokenLiteral() != fmt.Sprint(value) {
		t.Errorf("bo.TokenLiteral not %t. got=%s", value, bo.TokenLiteral())
		return false
	}
	return true
}
