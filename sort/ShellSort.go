package main

import "fmt"

func main() {
	a := []int{1, 54, 6, 3, 78, 34, 12, 45, 56, 100}

	shellSort(a)

	fmt.Println(a)

}

func shellSort(a []int) {
	d := len(a)

	for {
		d = d / 2

		for i := 0; i < d; i++ {
			for j := i + d; j < len(a); j += d {
				tmp := a[j]
				k := j - d
				for ; k > 0 && a[k] > tmp; k -= d {
					a[k+d] = a[k]
				}
				a[k+d] = tmp
			}
		}
		if d == 1 {
			break
		}
	}
}
