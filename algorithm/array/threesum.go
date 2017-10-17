package main

import (
	"sort"
	"fmt"
)
/**
 * http://www.lintcode.com/zh-cn/problem/3sum/

 给出一个有n个整数的数组S，在S中找到三个整数a, b, c，找到所有使得a + b + c = 0的三元组。
 注意事项

 在三元组(a, b, c)，要求a <= b <= c

 结果不能包含重复的三元组。

 样例
 如S = {-1 0 1 2 -1 -4}, 你需要返回的三元组集合的是：
 (-1, 0, 1)

 (-1, -1, 2)
 */

func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)

	for i := 0; i < len(nums) - 2; i++ {
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}
		a := nums[i]
		j, k := i + 1, len(nums) - 1
		for j < k {
			if j > i + 1 && nums[j] == nums[j - 1] {
				j++
				continue
			}
			b, c := nums[j], nums[k]
			sum := a + b + c
			if sum == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				k--
			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}
	}
	return res
}

func main() {
	a := []int{-1, 0, 1, 2, -1, -4}
	b := []int{-8, -2, -1, 0, 1, 2, 3, 6, 7, 8}
	c := []int{-2,0,0,2,2}
	fmt.Println(threeSum(a))
	fmt.Println(threeSum(b))
	fmt.Println(threeSum(c))
}
