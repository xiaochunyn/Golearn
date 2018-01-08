package main

import (
	"fmt"
)
/**
 * https://leetcode.com/problems/intersection-of-two-arrays/description/
 *
 * Given two arrays, write a function to compute their intersection.
 Example:
 Given nums1 = [1, 2, 2, 1], nums2 = [2, 2], return [2].

 Note:
 Each element in the result must be unique.
 The result can be in any order.
 */

// method 1, do not use other data structures
func intersection(a, b[]int) []int {
	var flag bool
	var res []int

	for _, i := range a {
		for _, j := range b{
			if i == j {
				flag = false
				for _, v := range res {
					if i == v {
						flag = true
					}
				}
				if !flag {
					res = append(res, i)
				}
			}
		}
	}
	return res
}

// method 2, binary search, time complexity is O(nlogn)
func intersection2(a, b []int) []int {
	var res []int
	for _, v := range a {
		if ok := binarySearch(b, v); ok {
			if exist := binarySearch(res, v); !exist {
				res = append(res, v)
			}
		}
	}
	return res
}

func binarySearch(a []int, target int) bool {
	start, end := 0, len(a) - 1
	for start <= end {
		mid := (start + end) / 2
		if target == a[mid] {
			 return true
		} else if target > a[mid] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return false
}

func main() {
	a := []int{1,2,2,2,3,4,3,5,6,}
	b := []int{1,2,3,2,3,6}
	fmt.Println(intersection(a, b))
	fmt.Println(intersection2(a, b))
}


