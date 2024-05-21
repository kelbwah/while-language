package evaluator

import (
	"fmt"
	"while/ast"
	"while/object"
    "strconv"
)

var (
    NULL  = &object.Null{} 
    TRUE  = &object.Boolean{Value: true}
    FALSE = &object.Boolean{Value: false}
)

func Eval(node ast.Node, env *object.Environment) object.Object {
    switch node := node.(type) {
    case *ast.Program:
        return evalStatements(node.Statements, env)
    case *ast.ExpressionStatement:
        return Eval(node.Expression, env)
    case *ast.BlockStatement:
        return evalStatements(node.Statements, env)
    case *ast.IfExpression:
        return evalIfExpression(node, env)
    case *ast.WhileExpression:
        return evalWhileExpression(node, env)
    case *ast.PrefixExpression:
        right := Eval(node.Right, env)
        if isError(right) {
            return right
        }
        return evalPrefixExpression(node.Operator, right)
    case *ast.InfixExpression:
        var left object.Object

        if node.Operator == "=" {
            left = &object.VariableName{Value: node.Left.String()}
        } else {
            left = Eval(node.Left, env)
            if isError(left) {
                return left
            }
        }

        right := Eval(node.Right, env)
        if isError(right) {
            return right
        }

        return evalInfixExpression(node.Operator, left, right, env)
    case *ast.IntegerLiteral:
        return &object.Integer{Value: node.Value}
    case *ast.Boolean:
        return nativeBoolToBooleanObject(node.Value)
    case *ast.Identifier:
        return evalIdentifier(node, env)
    default:
        return newError(fmt.Sprintf("Unknown token literal: %T\n", node.TokenLiteral()))
    }
}

func evalProgram(program *ast.Program, env *object.Environment) object.Object {
    var result object.Object

    for _, statement := range program.Statements {
        result = Eval(statement, env)

        switch result := result.(type) {
        case *object.Error:
            return result
        }
    }

    return result
}

func evalBlockStatement(block *ast.BlockStatement, env *object.Environment) object.Object {
    var result object.Object

    for _, statement := range block.Statements {
        result = Eval(statement, env)
        
        if result != nil {
            rt := result.Type()
            if rt == object.ERROR_OBJ {
                return result
            }
        }
    }

    return result
}

func evalStatements(statements []ast.Statement, env *object.Environment) object.Object {
    var result object.Object

    for _, statement := range statements {
        result = Eval(statement, env)
    }

    return result
}

func nativeBoolToBooleanObject(input bool) *object.Boolean {
    if input {
        return TRUE
    }

    return FALSE
}

func evalPrefixExpression(operator string, right object.Object) object.Object {
    switch operator {
    case "not":
        return evalNotOperatorExpression(right)
    case "-":
        return evalMinusPrefixOperatorExpression(right)
    default:
        return newError("unknown operator prefix: %s%s", operator, right.Type()) 
    }
}

func evalInfixExpression(operator string, left, right object.Object, env *object.Environment) object.Object {
    switch {
    case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
        return evalIntegerInfixExpression(operator, left, right)
    case operator == "=" && left.Type() != object.INTEGER_OBJ && left.Type() != object.BOOLEAN_OBJ && left.Type() != object.NULL_OBJ:
        env.Set(left.Inspect(), right)

        return right 
    case operator == "==":
        return nativeBoolToBooleanObject(left == right)
    case operator == "or":
        leftBool, _ := strconv.ParseBool(left.Inspect())
        rightBool, _ := strconv.ParseBool(right.Inspect())

        return nativeBoolToBooleanObject(leftBool || rightBool)
    case operator == "and":
        leftBool, _ := strconv.ParseBool(left.Inspect())
        rightBool, _ := strconv.ParseBool(right.Inspect())

        return nativeBoolToBooleanObject(leftBool && rightBool)
    default:
        return newError("unknown operator infix: %s %s %s",
            left.Type(), operator, right.Type())
    }
}

func evalIntegerInfixExpression(operator string, left, right object.Object) object.Object {
    leftVal := left.(*object.Integer).Value
    rightVal := right.(*object.Integer).Value
    
    switch operator {
        case "+":
            return &object.Integer{Value: leftVal + rightVal}
        case "-":
            return &object.Integer{Value: leftVal - rightVal}
        case "*": return &object.Integer{Value: leftVal * rightVal}
        case "<":
            return nativeBoolToBooleanObject(leftVal < rightVal)
        case "==":
            return nativeBoolToBooleanObject(leftVal == rightVal)
        case "or":
            return nativeBoolToBooleanObject(true)
        case "and":
            return nativeBoolToBooleanObject(true)
        default:
            return newError("unknown operator integer: %s %s %s",
               left.Type(), operator, right.Type())
    }
}

func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
    if right.Type() != object.INTEGER_OBJ {
        return newError("unknown operator: -%s", right.Type())
    }

    value := right.(*object.Integer).Value
    return &object.Integer{Value: -value}
}

func evalNotOperatorExpression(right object.Object) object.Object {
    switch right {
    case TRUE:
        return FALSE
    case FALSE:
        return TRUE
    case NULL:
        return TRUE
    default:
        return FALSE
    }
}

func evalIfExpression(ie *ast.IfExpression, env *object.Environment) object.Object {
    condition := Eval(ie.Condition, env)
    if isError(condition) {
        return condition
    }
    
    if isTruthy(condition) {
        return Eval(ie.Consequence, env)
    } else if ie.Alternative != nil {
        return Eval(ie.Alternative, env)
    } else {
        return newError("Invalid if expression")
    }
}

func evalWhileExpression(ie *ast.WhileExpression, env *object.Environment) object.Object {
    var output string
    condition := Eval(ie.Condition, env)
    
    if isError(condition) {
        return condition
    }
    
    for isTruthy(condition) {
        result := Eval(ie.Consequence, env)
        if isError(result) {
            return result 
        }
        if result != NULL {
            output += result.Inspect() + "\n"
        }
        condition = Eval(ie.Condition, env)
        if isError(condition) {
            return condition
        }
    } 

    return &object.StringResult{Value: output}
}

func evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {
    val, ok := env.Get(node.Value)
    if !ok {
        return newError("identifier not found: " + node.Value)
    }

    fmt.Printf("%s: %s\n", node.Value, val.Inspect())
    return val
}

func isTruthy(obj object.Object) bool {
    switch obj {
    case NULL:
        return false
    case TRUE:
        return true
    case FALSE:
        return false
    default:
        return true 
    }
}

func isError(obj object.Object) bool {
    if obj != nil {
        return obj.Type() == object.ERROR_OBJ
    }

    return false
}

func newError(format string, a ...interface{}) *object.Error {
    return &object.Error{ Message: fmt.Sprintf(format, a...) }
}
