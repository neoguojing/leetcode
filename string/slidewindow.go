package string

import "fmt"

// MinWindow 求s的最小窗口，包含t中的所有字符；窗口必须连续
// no 76
//  s = "ADOBECODEBANC", t = "ABC" result：BANC
func MinWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	if s == "" || t == "" {
		return ""
	}

	i, j := 0, 0
	tagetMap := map[byte]int{}

	for k := 0; k < len(t); k++ {
		if _, ok := tagetMap[t[k]]; !ok {
			tagetMap[t[k]] = 1
		} else {
			tagetMap[t[k]]++
		}
	}
	sMap := map[byte]int{}
	winSet := []int{}
	winCount := 0
	for j < len(s) {
		if _, ok := tagetMap[s[j]]; ok {
			if _, ok := sMap[s[j]]; !ok {
				sMap[s[j]] = 0
			}

			if sMap[s[j]] < tagetMap[s[j]] {
				sMap[s[j]]++
				winCount++
			} else {
				sMap[s[j]]++
			}
		}

		j++
		if winCount < len(t) {
			continue
		}

		for i < j {
			if _, ok := tagetMap[s[i]]; !ok {
				i++
			} else {
				if len(winSet) == 0 || j-i < winSet[0] {
					winSet = []int{j - i, i, j}
				}
				sMap[s[i]]--
				if sMap[s[i]] < tagetMap[s[i]] {
					winCount--
					i++
					break
				}
			}
		}

		fmt.Println(winSet)

	}

	return s[winSet[1]:winSet[2]]
}
