package utils

//Max ...
func Max(in ...int) int {
	if in == nil {
		return 0
	}

	if len(in) == 0 {
		return in[0]
	}

	max := in[0]
	for _, v := range in {
		if v > max {
			max = v
		}
	}

	return max
}

//Min ...
func Min(in ...int) int {
	if in == nil {
		return 0
	}

	if len(in) == 0 {
		return in[0]
	}

	min := in[0]
	for _, v := range in {
		if v < min {
			min = v
		}
	}

	return min
}

//Abs ...
func Abs(a, b int) int {
	if a-b > 0 {
		return a - b
	}
	return b - a
}
