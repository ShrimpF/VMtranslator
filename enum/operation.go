package enum

// Operation -- arithemtic op enum
type Operation int

const (
	// Add -- add a+b
	Add Operation = iota
	// Sub -- substract a-b
	Sub
	// Neg --  reverse the +/- sign
	Neg
	// Eq -- compare 2 nums.if a == b return true
	Eq
	// Gt --comapre 2 nums if a > b return true
	Gt
	// Lt -- comapre 2 nums if a < b return true
	Lt
	// And -- and operation
	And
	// Or -- or oparation
	Or
	// Not -- not opeartion
	Not
	// UndefinedOp -- undefined oppai
	UndefinedOp
)

func (op Operation) String() string {
	switch op {
	case Add:
		return "Add"
	case Sub:
		return "Sub"
	case Neg:
		return "Neg"
	case Eq:
		return "Eq"
	case Gt:
		return "Gt"
	case Lt:
		return "Lt"
	case And:
		return "And"
	case Or:
		return "Or"
	case Not:
		return "Not"
	default:
		return "Undefined"
	}
}
