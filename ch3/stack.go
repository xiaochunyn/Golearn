package stack

type Stack struct {
	i    int
	data []int
}

func (s *Stack) Push(i int) {
	s.data[i] = i
	s.i++
}

func (s *Stack) Pop() int {
	s.i--
	return s.data[s.i]
}
