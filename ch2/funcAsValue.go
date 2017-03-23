package main

import (
	"fmt"
)

func main() {
	a := func() {
		fmt.Println("Hello world")
	}
	a() //call a
	//fmt.Printf("\%T\n", a)
	fmt.Printf("%T\n", a)
}
