package game

const (
	nopOp = "nop"
	accOp = "acc"
	jmpOp = "jmp"
)

// Operations represent a list of instructions
type Operations []Operation

// Operation represents an instruction with its name, its value
//and a boolean set to true it is already seen
type Operation struct {
	name  string
	value int
	seen  bool
}

func NewOperations(len int) Operations {
	return make(Operations, len)
}

func NewOperation(name string, value int) Operation {
	return Operation{name, value, false}
}
