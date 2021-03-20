package targetsum

// compute the target sum by storing in a map the complement of the current value
// check for each value if the complement has been seen
// complexity time: O(n)
func  ComputeTargetSumWithPair(expenses []int, target int) int {
	complements := make(map[int]struct{})

	for _, value := range expenses {
		wantedComp := target - value
		complements[wantedComp] = struct{}{}

		if _, exist := complements[wantedComp]; exist {
			return wantedComp * value
		}
	}

	return 0
}

// Find the triplet that sums the target value
// complexity time: O(n²)
func ComputeTargetSumWithTriplet(expenses []int, target int) int {
	complements := make(map[int]struct{})

	// parse once all the values from the array
	for i := 0; i < len(expenses)-2; i++ {

		// parse the sub array from the array
		for j := i + 1; j < len(expenses); j++ {
			// check if the target value - the current value - the subarray value have already been seen ?
			if _, exist := complements[target-expenses[i]-expenses[j]]; exist {
				return (target - expenses[i] - expenses[j]) * expenses[i] * expenses[j]
			}
			// store the seen value
			complements[expenses[j]] = struct{}{}
		}
	}
	return 0
}


// Compute the target sum by parsing twice the integer array
// complexity time: O(n²)
func computeBrutallyTargetSumWithPair(expenses []int, target int) int {
	for i := 0; i < len(expenses); i++ {
		for j := i + 1; j < len(expenses); j++ {
			if i != j && expenses[i]+expenses[j] == target {
				return expenses[i] * expenses[j]
			}
		}
	}
	return -1
}
