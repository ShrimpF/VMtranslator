package parser

import (
	"strconv"
	"strings"

	"github.com/ShrimpF/vmTranslator/enum"
)

// Parser is basic struct to parse
type Parser struct {
	code string
}

// New create new Parser
func New() *Parser {
	return &Parser{}
}

// SetCode is change the code
func (p *Parser) SetCode(code string) {
	// remove comment
	index := strings.Index(code, "//")
	if index != -1 {
		code = code[:index]
	}

	p.code = code
}

// CommandType returns vm commands like C_Arithmetic,push pop,and so on.
func (p *Parser) CommandType() enum.Command {
	cmd := strings.Split(p.code, " ")
	switch cmd[0] {
	case "add":
		return enum.CArithmetic
	case "sub":
		return enum.CArithmetic
	case "neg":
		return enum.CArithmetic
	case "eq":
		return enum.CArithmetic
	case "gt":
		return enum.CArithmetic
	case "lt":
		return enum.CArithmetic
	case "and":
		return enum.CArithmetic
	case "or":
		return enum.CArithmetic
	case "not":
		return enum.CArithmetic
	case "push":
		return enum.CPush
	case "pop":
		return enum.CPop
	case "label":
		return enum.CLabel
	case "goto":
		return enum.CGoto
	case "if-goto":
		return enum.CIf
	case "function":
		return enum.CFunction
	case "call":
		return enum.CCall
	case "return":
		return enum.CReturn
	default:
		return enum.Undefined
	}
}

// Arg1 return the first argument of the code
// C_RETURN can't use this command
func (p *Parser) Arg1() string {
	code := strings.Split(p.code, " ")
	if p.CommandType() == enum.CArithmetic {
		return code[0]
	}
	return code[1]
}

// Arg2 return the second argument of the code
// Only CPush,CPop,Cfunc,Ccall can use this command
func (p *Parser) Arg2() int {
	code := strings.Split(p.code, " ")
	num, err := strconv.Atoi(code[2])
	if err != nil {
		panic(err)
	}
	return num
}

// Operation return arithmetic op like add sub,and or ...
func (p *Parser) Operation() enum.Operation {
	switch p.Arg1() {
	case "add":
		return enum.Add
	case "sub":
		return enum.Sub
	case "neg":
		return enum.Neg
	case "eq":
		return enum.Eq
	case "gt":
		return enum.Gt
	case "lt":
		return enum.Lt
	case "and":
		return enum.And
	case "or":
		return enum.Or
	case "not":
		return enum.Not
	default:
		return enum.UndefinedOp
	}
}
