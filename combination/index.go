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
