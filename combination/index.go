package combination

/*StringVectorMul ...
字符串向量相乘
*/
func StringVectorMul(a, b string) []string {
	ret := make([]string, 0)
	if a == "" {
		for i := range b {
			ret = append(ret, string(b[i]))
		}
	}

	if b == "" {
		for i := range a {
			ret = append(ret, string(a[i]))
		}
	}

	for i := range a {
		for j := range b {
			solu := string(a[i]) + string(b[j])
			ret = append(ret, solu)
		}
	}

	return ret
}

func StringVectorMul1(a, b []string) []string {
	ret := make([]string, 0)
	if len(a) == 0 {
		return b
	}

	if len(b) == 0 {
		return a
	}

	for _, v1 := range a {
		for _, v2 := range b {
			solu := v1 + v2
			ret = append(ret, solu)
		}
	}

	return ret
}
