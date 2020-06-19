package integer

/**
本文件用于边界考察
**/

//no 7
//整数反转
func reverse(x int32) int32 {
	var ret int32 = 0
	for x!=0 {
		mod := x%10
		x = x/10

		if ret > MAX_INT32/10 {
			return 0
		}

		if ret < MIN_INT32/10 {
			return 0
		}

		ret = ret * 10 + mod
	}

	return ret
}