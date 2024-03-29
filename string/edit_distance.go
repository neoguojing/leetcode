package string

import (
	"leetcode/utils"
)

/*MinDistance ...
no 72
假设序列S和T的长度分别为m和n, 两者的编辑距离表示为edit[m][n]. 则对序列进行操作时存在以下几种情况:

a, 当S和T的末尾字符相等时, 对末尾字符不需要进行上述定义操作中(亦即"编辑")的任何一个, 也就是不需要增加计数. 则满足条件: edit[m][n] = edit[m - 1][n - 1].

b, 当S和T的末尾字符不相等时, 则需要对两者之一的末尾进行编辑, 相应的计数会增加1.

b1, 对S或T的末尾进行修改, 以使之与T或S相等, 则此时edit[m][n] = edit[m - 1][n - 1] + 1;

b2, 删除S末尾的元素, 使S与T相等, 则此时edit[m][n] = edit[m - 1][n] + 1;

b3, 删除T末尾的元素, 使T与S相等, 则此时edit[m][n] = edit[m][n - 1] + 1;

b4, 在S的末尾添加T的尾元素, 使S和T相等, 则此时S的长度变为m+1, 但是此时S和T的末尾元素已经相等, 只需要比较S的前m个元素与T的前n-1个元素, 所以满足edit[m][n] = edit[m][n - 1] + 1;

b5, 在T的末尾添加S的尾元素, 使T和S相等, 此时的情况跟b4相同, 满足edit[m][n] = edit[m - 1][n] + 1;

c, 比较特殊的情况是, 当S为空时, edit[0][n] = n; 而当T为空时, edit[m][0] = m; 这个很好理解, 例如对于序列""和"abc", 则两者的最少操作为3, 即序列""进行3次插入操作, 或者序列"abc"进行3次删除操作.
*/
func MinDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)

	dp := initDP2DsForInt(m+1, n+1)
	for w2 := range word2 {
		dp[0][w2] = w2
	}
	dp[0][n] = n

	for w1 := range word1 {
		dp[w1][0] = w1
	}
	dp[m][0] = m

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if word1[i] == word2[j] {
				dp[i+1][j+1] = dp[i][j]
			} else {
				dp[i+1][j+1] = utils.Min(dp[i][j+1]+1, dp[i+1][j]+1, dp[i][j]+1)
			}
		}
	}

	return dp[m][n]
}

// LadderLength word ladder 即有一个字符不一样的单词，从beginword开始，从wordList选仅一个字符不一样的word，直到endWord命中
// 127
// BFS + editdistance
func LadderLength(beginWord string, endWord string, wordList []string) int {
	wordMap := map[string]bool{}
	for _, v := range wordList {
		wordMap[v] = false
	}

	if _, ok := wordMap[endWord]; !ok {
		return 0
	}

	q := utils.NewQueue()
	q.Push(beginWord)
	count := 1
	for !q.Empty() {
		qsize := q.Len()
		for i := 0; i < qsize; i++ {
			p := q.Pop().(string)
			if p == endWord {
				return count
			}

			for i := 0; i < len(p); i++ {
				strByte := []byte(p)
				for k := 0; k < 26; k++ {
					strByte[i] = byte('a' + k)
					newStr := string(strByte)
					if _, ok := wordMap[newStr]; ok {
						delete(wordMap, newStr)
						q.Push(newStr)
					}
				}

			}
		}

		count++
	}

	return 0
}

func findiff1(a string, words map[string]bool) []string {
	ret := []string{}
	for i := 0; i < len(a); i++ {
		strByte := []byte(a)
		for k := 0; k < 26; k++ {
			strByte[i] = byte('a' + k)
			newStr := string(strByte)
			if _, ok := words[newStr]; ok {
				ret = append(ret, newStr)
			}
		}
	}

	return ret
}

// FindLadders 同上，本题要求返回，所有的最短路径
// 126
func FindLadders(beginWord string, endWord string, wordList []string) [][]string {
	wordMap := map[string]bool{}
	for _, v := range wordList {
		wordMap[v] = false
	}

	ret := [][]string{}
	min := len(wordList)
	ladderHelper(1, &min, wordMap, beginWord, endWord, []string{beginWord}, &ret)
	return ret
}

func ladderHelper(level int, min *int, words map[string]bool, beginWord, endWord string, cur []string, ret *[][]string) {

	if beginWord == endWord {
		if *min > level {
			*min = level
			*ret = [][]string{}
			tmp := make([]string, len(cur))
			copy(tmp, cur)
			*ret = append(*ret, tmp)
		} else if level == *min {
			tmp := make([]string, len(cur))
			copy(tmp, cur)
			*ret = append(*ret, tmp)
		} else {

		}
		return
	}

	nerbs := findiff1(beginWord, words)
	if len(nerbs) == 0 {
		*min = 0
		*ret = [][]string{}
		return
	}

	for _, nerb := range nerbs {
		if !words[nerb] && level < *min {
			level++
			words[nerb] = true
			cur = append(cur, nerb)
			ladderHelper(level, min, words, nerb, endWord, cur, ret)
			cur = cur[:len(cur)-1]
			words[nerb] = false
			level--
		}

	}
}

// 433
