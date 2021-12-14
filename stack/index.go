package stack

type Stack []int

func (s Stack) Push(a int) {
	s = append(s, a)
}

func (s Stack) Pop() int {
	if len(s) > 0 {
		top := s[len(s)-1]
		s = s[:len(s)-1]
		return top
	}

	return -1
}

func (s Stack) Top() int {
	if len(s) > 0 {
		return s[len(s)-1]
	}
	return -1
}
