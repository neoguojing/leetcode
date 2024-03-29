package bits

import (
	"strconv"
)

// HammingWeight ...
// no 191
// 给一个数字，求出其二进制形式中 1 的个数。
// 以下，每次减一，会把借位的1计数
// 0111 0110   末尾为1，会把末尾的1计数
// 0110 0101  末尾非1，会把左边倒数第一个1计数
// 0101 0100
// 0100 0011   当只有1个1的情况下，以下方法会把值变为0
func HammingWeight(n int) int {
	count := 0
	for n != 0 {
		n &= (n - 1)
		count++
	}

	return count
}

//ReverseBits ...
// no 190
// 将一个 int 类型的数字，32 个 bit，进行倒置。
func ReverseBits(n int) int {
	ret := 0
	count := 0
	for count < 32 {
		ret = ret << 1
		ret = ret + n&1
		n = n >> 1
		count++
	}

	return ret
}

// AddBinary ...
//no 67
// 两个二进制数相加，返回结果
func AddBinary(a string, b string) string {
	m := len(a) - 1
	n := len(b) - 1

	ret := ""
	carry := 0

	bA := []rune(a)
	bB := []rune(b)
	for m >= 0 || n >= 0 {
		num1 := 0
		num2 := 0

		if m >= 0 {
			num1 = int(bA[m] - '0')
			m--
		}
		if n >= 0 {
			num2 = int(bB[n] - '0')
			n--
		}

		tmp := num1 + num2 + carry
		carry = tmp / 2
		cur := tmp % 2

		ret = strconv.Itoa(cur) + ret
	}

	if carry != 0 {
		ret = strconv.Itoa(carry) + ret
	}

	return ret
}

// Divide
// no 29
const MAX_INT = int32(^uint32(0) >> 1)
const MIN_INT = ^MAX_INT

func Divide(dividend int, divisor int) int {
	if divisor == 0 {
		return dividend
	}

	var count int64 = 0
	var sign int64 = 1
	var divd, divr int64 = int64(dividend), int64(divisor)
	if dividend < 0 && divisor > 0 {
		sign = -1
		divd = -divd
	} else if dividend > 0 && divisor < 0 {
		sign = -1
		divr = -divr
	} else if dividend < 0 && divisor < 0 {
		divr = -divr
		divd = -divd
	}

	if divd-divr < 0 {
		return 0
	}

	if divr == 1 {
		count = divd
	} else {
		for tmp := divr; divd-tmp >= 0; tmp += divr {
			count++
		}
	}

	if count*sign > int64(MAX_INT) {
		return int(MAX_INT)
	}

	if count*sign < int64(MIN_INT) {
		return int(MIN_INT)
	}
	return int(count * sign)
}

// SingleNumber3 o(n) o(1) 有两个出现一次，其他出现两次
// no 260
func SingleNumber3(nums []int) []int {
	mask := 0
	for i := 0; i < len(nums); i++ {
		mask ^= nums[i]
	}
	lowBit := mask & (-mask)

	x := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]&lowBit != 0 {
			x ^= nums[i]
		}
	}
	return []int{x, mask ^ x}
}

// SingleNumber2 o(n) o(1) 仅1个出现1次，其他出现3次
// no 137
func SingleNumber2(nums []int) int {
	// 出现3次，计数器为11（二进制），则需要两个int。分别表示计数器
	x1, x2, mask := 0, 0, 0
	for i := range nums {
		// x1为1，新来一个数，则x2需要进1
		x2 ^= x1 & nums[i]

		// 相同为0，不同为1
		x1 ^= nums[i]
		// mask作用，是当相同的数累计三次之后，清零计数器；值为0或1
		mask = ^(x1 & x2)
		// 清理计数器
		x1 = x1 & mask
		x2 = x2 & mask
	}
	return x1
}

// SingleNumber o(n) o(1) 仅1个出现1次，其他出现2次
// no 136
// 亦或
func SingleNumber(nums []int) int {
	ret := nums[0]
	for i := 1; i < len(nums); i++ {
		ret ^= nums[i]
	}
	return ret
}
