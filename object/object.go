package object

import "fmt"

type ObjectType string

const (
	IntegerTypeObj ObjectType = "INTEGER"
	BooleanTypeObj ObjectType = "BOOLEAN"
	NullTypeObj    ObjectType = "NULL"
	ReturnTypeObj  ObjectType = "RETURN"
	ErrorObj       ObjectType = "ERROR"
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