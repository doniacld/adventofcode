package game

// ComputeFixedAcc returns the accumulator value
// game ends at the last instruction
// a jmp is replaced by a nop or the other way around
func (ops Operations) ComputeFixedAcc() int {
	for i := 0; i < len(ops)-1; i++ {

		o := ops[i]
		tmpOps := ops.DuplicateOps()

		// switch on the nop or jmp to check if the replace is successful
		if o.value > 1 {
			continue
		}

		switch o.name {
		case nopOp:
			tmpOps.replace(i, jmpOp)
		case jmpOp:
			tmpOps.replace(i, nopOp)
		}

		// if the game ends at the last instruction, return accCounter
		v, accCounter := tmpOps.computeToTheLastOp()
		if v {
			return accCounter
		}
	}
	// should not happened
	return 0
}

func (ops Operations) DuplicateOps() Operations {
	// store the tmp ops for modification
	tmpOps := NewOperations(len(ops))
	copy(tmpOps, ops)
	return tmpOps
}

// replace the current instruction by the new wanted one
func (ops Operations) replace(idx int, newOp string) {
	ops[idx].name = newOp
}

// computeToTheLastOp calculates the accumulator counter depending on operations
// returns false if the  operation is already seen or the index is not in the range anymore
func (ops Operations) computeToTheLastOp() (bool, int) {
	currentIdx, nextIdx, accCounter := 0, 0, 0

	o := ops[currentIdx]
	for ; nextIdx != len(ops)-1; o = ops[currentIdx] {
		switch o.name {
		case nopOp:
			nextIdx = ops.nops(currentIdx)
			break
		case accOp:
			nextIdx, accCounter = ops.accs(currentIdx, accCounter)
			break
		case jmpOp:
			nextIdx = ops.jumps(currentIdx)
			break
		}

		// operation already seen, this is not normal
		if ops[nextIdx].seen {
			return false, accCounter
		}

		// nextIdx not in the range anymore
		if nextIdx < 0 || nextIdx > len(ops)-1 {
			return false, accCounter
		}

		// let's continue, Operation is seen and save the index
		ops[nextIdx].seen = true
		currentIdx = nextIdx
	}

	// compute the last value
	if ops[nextIdx].name == accOp {
		_, accCounter = ops.accs(nextIdx, accCounter)
	}
	return true, accCounter
}

// nop for no Operation does nothing, go to the next instruction
func (ops Operations) nops(idx int) int {
	return idx + 1
}

// acc for accumulator, increase the current value of the acc counter
// and go to the next instruction
func (ops Operations) accs(idx int, acc int) (int, int) {
	if idx == len(ops)-1 {
		return idx, ops[idx].value + acc
	}
	return idx + 1, ops[idx].value + acc
}

// jump add the offset to the current index and go to this instruction
func (ops Operations) jumps(idx int) int {
	return idx + ops[idx].value
}
