package main

import "fmt"

func main() {
	//a := []int{1, 54, 6, 3, 78, 34, 12, 45, 56, 100}
	a := []int{4, 2, 5, 7, 3, 9, 6, 8}
	shellSort(a)

	fmt.Println(a)

}

func shellSort(a []int) {
	d := len(a)

	for gap := d / 2; gap > 0; gap /= 2 {

		for i := gap; i < d; i++ {
			for j := i - gap; j >= 0 && a[j] > a[j+gap]; j -= gap {
				a[j], a[j+gap] = a[j+gap], a[j]
			}
		}
	}
}
