package main

import "fmt"

type stack struct {
	i    int
	data [10]int
}

func (s *stack) push(val int) {

	s.data[s.i] = val
	s.i++
}

func (s *stack) pop() int {
	if s.i < 0 || s.i >= 10 {
		return -1
	}
	s.i--
	return s.data[s.i]
}

func main() {
	var s stack
	s.push(23)
	s.push(5)
	fmt.Println(s)
	s.pop()
	fmt.Println(s)
}
