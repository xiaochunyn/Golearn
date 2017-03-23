package main

import "fmt"

func main() {
	astring := "This is a string."
	fmt.Printf("%v\n", astring)

	byteslice := []byte(astring)
	fmt.Printf("%v\n", byteslice)

	runeslice := []rune(astring)
	fmt.Printf("%v\n", runeslice)

	var i int = 2
	var f float32 = 3.8

	a := int8(i)
	fmt.Println(a)

	b := int(f)
	fmt.Println(b)

	c := float32(i)
	fmt.Println(c)
}
