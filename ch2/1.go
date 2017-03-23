package main

import "fmt"

func main() {
	a := 1
	rec(a)
}

func rec(i int) {
	if i == 10 {
		return
	}
	rec(i + 1)
	fmt.Printf("Val is : %d\n", i)
}
