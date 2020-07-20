package bits

import "strconv"

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
