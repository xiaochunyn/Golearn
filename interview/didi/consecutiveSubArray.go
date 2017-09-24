package main

import "fmt"

/**
给定一个数组 a = [1 2 3 5 8 7 8 4],
输出连续的字串，如 [1 2 3] [7 8]
*/

func main() {
	var n int
	fmt.Scanln(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	var index, pivot int
	for i := 1; i < n; i++ {
		if a[i] == a[index]+1 {
			index++
		} else {
			if index != pivot {
				fmt.Println(a[pivot:i])
			}
			index = i
			pivot = i
		}
	}
}

/**
8
1 2 3 5 8 7 8 4
*/
