package string

// IsInterleave ...
// no 97
// 在两个字符串 s1 和 s2 中依次取字母，问是否可以组成 S3。什么意思呢？比如 s1 = abc , s2 = de，s3 = abdce。
// len(s1) + len(s2) == len(s3)
//回溯
func IsInterleave(s1, s2, s3 string) bool {
	return isInterleave(s1, s2, s3, 0, 0, 0)
}

func isInterleave(s1, s2, s3 string, i, j, k int) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	if i == len(s1) && j == len(s2) && k == len(s3) {
		return true
	}

	if i == len(s1) {
		for j < len(s2) {
			if s2[j] != s3[k] {
				return false
			}
			j++
			k++
		}
		return true
	}

	if j == len(s2) {
		for i < len(s1) {
			if s1[i] != s3[k] {
				return false
			}
			i++
			k++
		}
		return true
	}

	if s1[i] == s3[k] {
		if isInterleave(s1, s2, s3, i+1, j, k+1) {
			return true
		}
	}

	if s2[j] == s3[k] {
		if isInterleave(s1, s2, s3, i, j+1, k+1) {
			return true
		}
	}

	return false
}
