package enum

// Command is enum of commands
type Command int

const (
	// CArithmetic is arithmetic operations like add sub ...
	CArithmetic Command = iota
	// CPush means push command
	CPush
	// CPop means pop command
	CPop
	// CLabel means label
	CLabel
	// CGoto means jump
	CGoto
	// CIf means if
	CIf
	// CFunction is
	CFunction
	// CReturn is
	CReturn
	// CCall is
	CCall
	// Undefined is
	Undefined
)

func (cmd Command) String() string {
	switch cmd {
	case CArithmetic:
		return "C_ARITHMETIC"
	case CPush:
		return "C_PUSH"
	case CPop:
		return "C_POP"
	case CLabel:
		return "C_Label"
	case CGoto:
		return "C_GOTO"
	case CIf:
		return "C_IF"
	case CFunction:
		return "C_FUNCTION"
	case CReturn:
		return "C_RETURN"
	case CCall:
		return "C_CALL"
	default:
		return "Undefined"
	}
}
