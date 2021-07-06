package game

// ComputeFixedAcc returns the accumulator value
// game ends at the last instruction
// a jmp is replaced by a nop or the other way around
func (ops Operations) ComputeFixedAcc() int {
	for i := 0; i < len(ops)-1; i++ {

		o := ops[i]
		tmpOps := ops.ResetOps()

		// switch on the nop or jmp to check if the replace is successful
		switch o.name {
		case nopOp:
			if o.value == 0 || o.value == 1 {
			} else {
				tmpOps.replace(i, jmpOp)
				break
			}
		case jmpOp:
			if o.value == 1 {
			} else {
				tmpOps.replace(i, nopOp)
				break
			}
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

func (ops Operations) ResetOps() Operations {
	// store the tmp ops for modification
	tmpOps := make(Operations, len(ops))
	copy(tmpOps, ops)
	return tmpOps
}

// replace the current instruction by the new wanted one
func (ops Operations) replace(idx int, newOp string) {
	ops[idx].name = newOp
}

// computeToTheLastOp calculates the accumulator counter depending on operations
// exit false if the the operation is already seen or the index is not anymore in the range
func (ops Operations) computeToTheLastOp() (bool, int) {
	i, accCounter := 0, 0

	for true {
		o := ops[i]
		idx := 0

		switch o.name {
		case nopOp:
			idx = ops.nops(i)
			break
		case accOp:
			idx, accCounter = ops.accs(i, accCounter)
			break
		case jmpOp:
			idx = ops.jumps(i)
			break
		}

		//  positive case
		if idx == len(ops)-1 {
			// compute the last value
			if ops[idx].name == accOp {
				_, accCounter = ops.accs(idx, accCounter)
			}
			return true, accCounter
		}

		// Operation already seen, this is not normal
		if ops[idx].seen {
			return false, accCounter
		}

		//idx not in the range anymore
		if idx < 0 || idx > len(ops)-1 {
			return false, accCounter
		}

		// let's continue, Operation is seen and save the index
		ops[idx].seen = true
		i = idx
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
