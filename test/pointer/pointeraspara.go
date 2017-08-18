package main

import "fmt"

func main() {
	//定义局部变量
	var a int = 10
	var b int = 20

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	swap(a, b)

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	swaps(&a, &b)
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
}

func swap(a, b int) {
	a, b = b, a
}

func swaps(a, b *int) {
	*a, *b = *b, *a
}
