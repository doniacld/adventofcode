package game

// ComputeCorruptedAcc computes the accumulator in the corrupted game
// ends when an Operation is seen twice
func (ops Operations) ComputeCorruptedAcc() (idx int, accCounter int) {

	o := ops[idx]
	for ; !o.seen; o = ops[idx] {
		switch o.name {
		case nopOp:
			idx = ops.nop(idx)
		case accOp:
			idx, accCounter = ops.acc(idx, accCounter)
		case jmpOp:
			idx = ops.jump(idx)
		}
	}
	return idx, accCounter
}

// nop for no Operation does nothing, go to the next instruction
func (ops Operations) nop(idx int) int {
	ops[idx].seen = true
	return (idx + 1) % len(ops)
}

// acc for accumulator, increase the current value of the acc counter
// and go to the next instruction
func (ops Operations) acc(idx int, acc int) (int, int) {
	ops[idx].seen = true
	return (idx + 1) % len(ops), ops[idx].value + acc
}

// jump add the offset to the current index and go to this instruction
func (ops Operations) jump(idx int) int {
	ops[idx].seen = true
	return (idx + ops[idx].value) % len(ops)
}
