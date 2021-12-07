package utils

func IntMax(a ...int) int {
	max := a[0]
	for _, i := range a {
		if i > max {
			max = i
		}
	}
	return max
}

func IntMin(a ...int) int {
	min := a[0]
	for _, i := range a {
		if i < min {
			min = i
		}
	}
	return min
}

func IntAbs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
