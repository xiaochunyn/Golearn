package main

import "fmt"

func main() {
	s := []float64{1.1, 1.2, 3.4, 4.5, 23, 23.2}

	var result float64
	length := len(s)
	for _, v := range s {
		result += v
	}
	fmt.Println(result / float64(length))
}
