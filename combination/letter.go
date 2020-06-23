package combination

import (
	"leetcode/utils"
)

/*LetterCombinations ...
排列组合

17 给一串数字，每个数可以代表数字键下的几个字母，返回这些数字下的字母的所有组成可能
1.相乘
2.队列
3.迭代
*/
func LetterCombinations(digits string) []string {

	var digitLetter = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	strings := make([]string, 0)
	for i := range digits {
		index := digits[i] - '0'
		strings = append(strings, digitLetter[index])
	}

	if len(strings) == 0 {
		return nil
	}

	if len(strings) == 1 {
		return StringVectorMul("", strings[0])
	}

	ret := StringVectorMul(strings[0], strings[1])
	for i := 2; i < len(strings); i++ {
		var tmp []string
		for _, v := range ret {
			tmp = append(tmp, StringVectorMul(v, strings[i])...)
		}
		ret = tmp
	}

	return ret

}

func LetterCombinationsWithQueue(digits string) []string {
	var digitLetter = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	strings := make([]string, 0)
	for i := range digits {
		index := digits[i] - '0'
		strings = append(strings, digitLetter[index])
	}

	if len(strings) == 0 {
		return nil
	}

	if len(strings) == 1 {
		return StringVectorMul("", strings[0])
	}

	q := utils.NewQueue()
	ret := StringVectorMul(strings[0], strings[1])
	for _, v := range ret {
		q.Push(v)
	}

	for i := 2; i < len(strings); i++ {
		var tmp []string
		for !q.Empty() {
			a := q.Pop().(string)
			tmp = append(tmp, StringVectorMul(a, strings[i])...)
		}
		for _, v := range tmp {
			q.Push(v)
		}
	}

	ret = make([]string, 0)
	for !q.Empty() {
		ret = append(ret, q.Pop().(string))
	}
	return ret
}
