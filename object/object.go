package object

import (
    "fmt"
)

type ObjectType string

const (
    INTEGER_OBJ = "INTEGER"
    BOOLEAN_OBJ = "BOOLEAN"
    NULL_OBJ = "NULL"
    ERROR_OBJ = "ERROR"
    VARIABLE_NAME_OBJECT = "VAR_NAME"
)

type Object interface {
    Type() ObjectType
    Inspect() string
}

type Integer struct {
    Value int64
}

type Boolean struct {
    Value bool 
}

type VariableName struct {
    Value string
}

type Error struct {
    Message string 
}

type Null struct {}

func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

func (v *VariableName) Inspect() string { return fmt.Sprintf("%s", v.Value) }
func (v *VariableName) Type() ObjectType { return VARIABLE_NAME_OBJECT }

func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }
func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }

func (n *Null) Inspect() string { return "null" }
func (n *Null) Type() ObjectType { return NULL_OBJ }

func (e *Error) Inspect() string { return fmt.Sprintf("Error: %s", e.Message) }
func (e *Error) Type() ObjectType { return ERROR_OBJ }
