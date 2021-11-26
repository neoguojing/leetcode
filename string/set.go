package string

// PartitionLabels 标签分类：将字符串分为多个集合(不能打乱顺序)，要求每个字符串只能出现在一个集合中，输出每个集合字符的数量
// no 763
func PartitionLabels(s string) []int {
	sets := make([][]int, 0)
	setMap := make(map[string]int)
	for i := 0; i < len(s); i++ {
		idx, ok := setMap[string(s[i])]
		if !ok {
			sets = append(sets, []int{i, i})
			setMap[string(s[i])] = len(sets) - 1
			// idx = len(sets) - 1
		} else {
			sets[idx][1] = i
		}
	}
	// fmt.Println(sets)
	// fmt.Println(setMap)
	ret := [][]int{sets[0]}
	for i := 1; i < len(sets); i++ {
		set := getUnionSet(ret[len(ret)-1], sets[i])
		if set == nil {
			ret = append(ret, sets[i])
		} else {
			ret[len(ret)-1] = set
		}
	}
	// fmt.Println(ret)

	result := make([]int, len(ret))
	for i := 0; i < len(ret); i++ {
		result[i] = ret[i][1] - ret[i][0] + 1
	}
	return result
}

func getUnionSet(a, b []int) []int {
	if a[0] > b[1] || a[1] < b[0] {
		return nil
	}

	l := a[0]
	if l > b[0] {
		l = b[0]
	}

	r := a[1]
	if r < b[1] {
		r = b[1]
	}

	return []int{l, r}
}
