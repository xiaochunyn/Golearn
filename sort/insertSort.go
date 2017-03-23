package main

import "fmt"

func main() {
	a := []int{57, 68, 59, 52}

	insertSort(a)

	fmt.Println(a)
}

func insertSort(a []int) []int {
	for i := 1; i < len(a); i++ {
		tmp := a[i]
		var j int
		//		for j = i - 1; j >= 0 && tmp < a[j]; j-- {
		//			a[j+1] = a[j]
		//		}

		for j = i - 1; j >= 0; j-- {
			if tmp < a[j] {
				a[j+1] = a[j]
			} else {
				break
			}

		}
		a[j+1] = tmp
	}
	return a
}
