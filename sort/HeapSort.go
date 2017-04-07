package main

import "fmt"

//import "math"

func main() {
	a := []int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1, 0}
	fmt.Printf("%v \n", a)
	sort(a)
	fmt.Printf("%v \n", a)
}

func swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}

func buildMaxHeap(a []int) {

	iParent := len(a)/2 - 1
	for i := iParent; i >= 0; i-- {
		maxHeapify(a, iParent, len(a))
	}
}

func maxHeapify(a []int, index, heapsize int) {
	var iMax, iLeft, iRight int

	for {
		iMax = index
		iLeft = 2*index + 1
		iRight = 2 * (index + 1)

		if iLeft < heapsize && a[iLeft] > a[iMax] {
			iMax = iLeft
		}

		if iRight < heapsize && a[iRight] > a[iMax] {
			iMax = iRight
		}

		if index != iMax {
			swap(a, iMax, index)
			index = iMax
		} else {
			break
		}
	}
}

func sort(a []int) {
	buildMaxHeap(a)
	for i := len(a) - 1; i >= 0; i-- {
		swap(a, 0, i)
		maxHeapify(a, 0, i)
	}
}
