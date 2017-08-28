package main

import "fmt"

func main() {
	a := [4]int{}
	a[0] = 1
	a[1] = 2
	fmt.Println(a)

	a[3] = 4
	fmt.Println(a)
}
