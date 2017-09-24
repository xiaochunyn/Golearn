package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	start, end := 0, len(a)-1
	for i := 0; i < n; i++ {
		if a[start] > 0 {
			start++
		}
		if a[end] < 0 {
			end--
		}
		if a[start] < 0 && a[end] > 0 {
			a[start], a[end] = a[end], a[start]
		}
		fmt.Println(a)
	}
	fmt.Println(a)
}

/**
6
-1 2 3 -2 4 -5
*/
