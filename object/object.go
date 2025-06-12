package object

import "fmt"

type ObjectType string

const (
	integerObj ObjectType = "INTEGER"
	booleanObj ObjectType = "BOOLEAN"
	nullObj    ObjectType = "NULL"
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
	return integerObj
}

type Boolean struct {
	Value bool
}

func (b *Boolean) Inspect() string {
	return fmt.Sprint(b.Value)
}

func (b *Boolean) Type() ObjectType {
	return booleanObj
}

type Null struct{}

func (n *Null) Inspect() string {
	return "null"
}

func (n *Null) Type() ObjectType {
	return nullObj
}
