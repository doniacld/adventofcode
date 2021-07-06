package game

// ComputeCorruptedAcc computes the accumulator in the corrupted game
// ends when an Operation is seen twice
func (ops Operations) ComputeCorruptedAcc() (int, int) {
	idx, accCounter := 0, 0
	for true {
		o := ops[idx]
		// op already seen, game over
		if o.seen {
			return idx, accCounter
		}

		// let's see which Operation to apply
		switch o.name {
		case nopOp:
			idx = ops.nop(idx)
			break
		case accOp:
			idx, accCounter = ops.acc(idx, accCounter)
			break
		case jmpOp:
			idx = ops.jump(idx)
			break
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
	if ops[idx].value < 0 {
		return idx - (abs(ops[idx].value) % len(ops))
	}
	return (idx + ops[idx].value) % len(ops)
}

// abs returns the absolute value of x.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
