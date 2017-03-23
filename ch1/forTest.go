package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("*******************************")

	// for j := 0; j < 10; j++ {
	// Here:
	// 	fmt.Println(j)
	// 	j++
	// 	goto Here
	// }
	j := 0
Loop:

	if j < 10 {
		fmt.Println(j)
		j++
		goto Loop
	}

}
