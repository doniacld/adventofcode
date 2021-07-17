package cypher

// FindSum browses the list of numbers until the retrieved value in part 1 and
// finds a list of contiguous values that sum up make this target value.
// It returns the sum of the min and max of this list.
func (c Cypher) FindSum(idx int) int {
	target := c.Numbers[idx]

	sum := 0
	tmpArray := make([]int, 0)
	for i := 0; i < idx-1; i++ {

		// goal reached
		if sum == target {
			return sumMinMax(tmpArray)
		}

		// delete first value from the array until
		// the sum is below the target
		for sum > target {
			sum -= tmpArray[0]
			s := tmpArray[1:]
			tmpArray = s
		}

		// add the next value in the array
		if sum < target {
			tmpArray = append(tmpArray, c.Numbers[i])
			sum += c.Numbers[i]
		}
	}
	return 0
}

// sumMinMax returns the sum of the min and max values from a given array
func sumMinMax(sum []int) int {
	min, max := sum[0], 0
	for _, v := range sum {

		if v < min {
			min = v
		}

		if v > max {
			max = v
		}
	}
	return min + max
}
