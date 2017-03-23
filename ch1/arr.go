package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5}
	s1 := a[2:4]
	s2 := a[1:5]
	s3 := a[2:]
	s4 := a[:5]
	s5 := s2[:]

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)

	ss := append(s5, 6)
	sa := append(s1, 6)
	fmt.Println(ss)
	fmt.Println(sa)

	fmt.Println(".......................")
	slice1 := append(s5, 7, 8, 9)
	slic32 := append(s1, s2...)
	fmt.Println(slice1)
	fmt.Println(slic32)

	var aArr = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	var aSlice = make([]int, 6)
	n1 := copy(aSlice, aArr[0:])
	fmt.Println(n1)
	fmt.Println(aSlice)

	n2 := copy(aSlice, aArr[2:])
	fmt.Println(n2)
	fmt.Println(aSlice)
}
