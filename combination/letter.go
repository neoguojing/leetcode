package combination

import (
	"leetcode/utils"
)

/*LetterCombinations ...
排列组合

no 17 给一串数字，每个数可以代表数字键下的几个字母，返回这些数字下的字母的所有组成可能
1.相乘
2.队列
3.迭代
*/
func LetterCombinations(digits string) []string {

	digitLetter := map[int][]string{
		0: nil,
		1: nil,
		2: []string{"a", "b", "c"},
		3: []string{"d", "e", "f"},
		4: []string{"g", "h", "i"},
		5: []string{"j", "k", "l"},
		6: []string{"m", "n", "o"},
		7: []string{"p", "q", "r", "s"},
		8: []string{"t", "u", "v"},
		9: []string{"w", "x", "y", "z"},
	}
	strings := make([][]string, len(digits))
	for i := range digits {
		index := int(digits[i] - '0')
		strings[i] = digitLetter[index]
	}

	if len(strings) == 0 {
		return nil
	}

	ret := letterCombinations(strings)
	return ret

}

func letterCombinations(in [][]string) []string {
	if len(in) == 0 {
		return []string{}
	}

	if len(in) == 1 {
		return in[0]
	}

	ret := StringVectorMul1(in[0], letterCombinations(in[1:]))
	return ret
}

//LetterCombinationsWithQueue ...
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

//GenerateParenthesis
// no 22 生成圆括号的组合，必须是合法的组合
/*
Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]
*/

const left = "("
const right = ")"

func GenerateParenthesis(n int) []string {

	if n == 1 {
		return []string{"()"}
	}
	var targets = []string{}
	generateParenthesis(0, 0, n, "", &targets)
	return targets
}

func generateParenthesis(l, r, n int, ret string, targets *[]string) {
	if l < r || l > n || r > n {
		return
	}

	if len(ret) == 2*n {
		*targets = append(*targets, ret)
		return
	}

	ret += left
	l++
	generateParenthesis(l, r, n, ret, targets)
	if l-r > 1 {
		l--
		ret = ret[:len(ret)-1]
		ret += right
		r++
		generateParenthesis(l, r, n, ret, targets)
	}

}
