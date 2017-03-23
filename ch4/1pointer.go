package main

import "fmt"

func main() {
	var p *int
	fmt.Println(p)

	var i int
	p = &i
	fmt.Println(p)

	i = 3
	fmt.Println(*p)

	*p = 8
	fmt.Print(i)
}
