package main

import "fmt"

func main() {
	func1(1, 2, 3, 4, 5)
}
func func1(arg ...int) {
	func2(arg[:2])
	func2(arg[2:])

	for _, i := range arg {
		fmt.Print(i)

	}
	fmt.Println()
}

func func2(arg ...[]int) {
	for _, i := range arg {
		fmt.Print(i)
	}
	fmt.Println()
}

func func2(arg ...interface{}) {
	for _, i := range arg {
		fmt.Print(i)
	}
	fmt.Println()
}
