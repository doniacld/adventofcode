package depth

func ComputeIncreaseDepth(m []int) int {
	cnt := 0

	for i := 1; i < len(m); i++ {
		if m[i] > m[i-1] {
			cnt++
		}
	}
	return cnt
}

func ComputeIncreaseDepthSlidingWindow(m []int) int {
	cnt := 0
	sum := 0
	for i := 1; i < len(m)-2; i++ {
		s := m[i] + m[i+1] + m[i+2]
		if s > sum {
			cnt++
		}
		sum = s
	}

	return cnt
}
