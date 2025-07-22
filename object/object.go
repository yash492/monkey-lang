package object

import (
	"fmt"
	"monkey/ast"
	"strings"
)

type ObjectType string

const (
	IntegerTypeObj ObjectType = "INTEGER"
	BooleanTypeObj ObjectType = "BOOLEAN"
	NullTypeObj    ObjectType = "NULL"
	ReturnTypeObj  ObjectType = "RETURN"
	ErrorObj       ObjectType = "ERROR"
	FunctionObj    ObjectType = "FUNCTION"
	StringObj      ObjectType = "STRING"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string {
	return fmt.Sprint(i.Value)
}

func (i *Integer) Type() ObjectType {
	return IntegerTypeObj
}

type Boolean struct {
	Value bool
}

func (b *Boolean) Inspect() string {
	return fmt.Sprint(b.Value)
}

func (b *Boolean) Type() ObjectType {
	return BooleanTypeObj
}

type Null struct{}

func (n *Null) Inspect() string {
	return "null"
}

func (n *Null) Type() ObjectType {
	return NullTypeObj
}

type Return struct {
	Value Object
}

func (r *Return) Inspect() string {
	return r.Value.Inspect()
}

func (r *Return) Type() ObjectType {
	return ReturnTypeObj
}

type Error struct {
	Message string
}

func (e *Error) Type() ObjectType {
	return ErrorObj
}

func (e *Error) Inspect() string {
	return "ERROR: " + e.Message
}

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() ObjectType {
	return FunctionObj
}

func (f *Function) Inspect() string {
	var out strings.Builder
	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")
	return out.String()
}

type String struct {
	Value string
}

func (s *String) Inspect() string {
	return s.Value
}

func (s *String) Type() ObjectType {
	return StringObj
}