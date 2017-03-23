package main

import "fmt"

//import "os"

func main() {
	fmt.Println("Hello world. ")

	// for i := 0; i < len(os.Args); i++ {
	// 	fmt.Println(os.Args[i])
	// }

	var a int
	var b int
	a = 15
	b = 16

	c := 12
	d := 13

	fmt.Println(a + b)
	fmt.Printf("c + d = %d \r\n", c+d)

	var s string
	s = "I Love "
	t := "You "
	fmt.Printf("I want to say : %s ", s+t)

	const (
		m = iota
		n = iota
	)
	fmt.Println(n)
	fmt.Println(m + n)

	fmt.Println("............................")

	var aa string
	aa = "aaaa"
	fmt.Println(&aa)

	aa = "bbbb"
	fmt.Println(&aa)

	cc := []rune(aa)
	cc[0] = 'a'
	ab := string(cc)
	fmt.Println(&ab)

	fmt.Println(".............................")

	if bo := false; bo == true {
		fmt.Println("bo is true")
	} else {
		fmt.Println("bo is false")
	}

	fmt.Println("............................")
	val := "a"
	switch val {
	case "a":
		fmt.Println("aaaaa")
	case "b":
		fmt.Println("bbbbb")
	default:
		fmt.Println("This is default. ")
	}
}
