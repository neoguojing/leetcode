package string

// RomanToInt 罗马数字转数字
// 13
func RomanToInt(s string) int {
	r2i := map[string]int{
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}

	ret := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && (string(s[i]) == "I" || string(s[i]) == "X" || string(s[i]) == "C") {
			if v, ok := r2i[s[i:i+2]]; ok {
				ret += v
				i++
			} else {
				ret += r2i[string(s[i])]
			}

		} else {
			ret += r2i[string(s[i])]
		}
	}

	return ret
}

// 从后往前遍历
func RomanToInt2(s string) int {
	r2i := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	ret := r2i[string(s[len(s)-1])]
	for i := len(s) - 2; i >= 0; i-- {
		if r2i[string(s[i])] < r2i[string(s[i+1])] {
			ret -= r2i[string(s[i])]
		} else {
			ret += r2i[string(s[i])]
		}
	}

	return ret
}

// 12
func IntToRoman(num int) string {
	return ""
}
