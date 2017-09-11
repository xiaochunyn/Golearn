package main

import (
	"fmt"
)

func main() {
	//var name string
	//var age int
	//n, err := fmt.Scanf("%s %d", &name, &age)
	//
	//fmt.Scanln(&name, &age)
	//fmt.Println(name, age)
	//fmt.Println(n, err, name, age)

	var a [5]int
	for i := 0; i < 5; i++ {
		fmt.Scanln(&a[i])
	}
	fmt.Printf("%v", a)
}
