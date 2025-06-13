package object

import "fmt"

type ObjectType string

const (
	IntegerTypeObj ObjectType = "INTEGER"
	BooleanTypeObj ObjectType = "BOOLEAN"
	NullTypeObj    ObjectType = "NULL"
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
