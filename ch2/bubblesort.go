package main

import "fmt"

func main() {
	a := []int{5, -1, 0, 12, 3, 5}
	fmt.Printf("Unsorted %v \n", a)
	bubblesort(a)
	fmt.Printf("Sorted %v \n", a)

	m, n := 1, 2
	exchange(m, n)
	//	x, y := exchange(m, n)
	//	fmt.Printf("x = %v, y = %v \n", x, y)

	fmt.Printf("m = %v, n = %v \n", m, n)
}

func exchange(a, b int) (x, y int) {
	fmt.Printf("....  a %v, b: %v \n", &a, &b)
	x, y = b, a
	return x, y
}

func bubblesort(a []int) {
	fmt.Printf("Address of a : %v \n", &a)

	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
				//				tmp := a[i]
				//				a[i] = a[j]
				//				a[j] = tmp
			}
		}
	}
}
