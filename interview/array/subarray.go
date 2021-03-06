package main

import (
	"fmt"
)

/**
Leetcode 53. Maximum Subarray
Find the contiguous subarray within an array (containing at least one number) which has the largest sum.

For example, given the array [-2,1,-3,4,-1,2,1,-5,4],
the contiguous subarray [4,-1,2,1] has the largest sum = 6.

滴滴2017秋招
[编程题] 连续最大和

时间限制：1秒
空间限制：32768K
一个数组有 N 个元素，求连续子数组的最大和。 例如：[-1,2,1]，和最大的连续子数组为[2,1]，其和为 3
输入描述:
输入为两行。
第一行一个整数n(1 <= n <= 100000)，表示一共有n个元素
第二行为n个数，即每个元素,每个整数都在32位int范围内。以空格分隔。
输出描述:
所有连续子数组中和最大的值。
输入例子1:
3
-1 2 1
输出例子1:
3
*/

func main() {
	var n int
	fmt.Scanln(&n)
	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	fmt.Println(maxSubArray(a))
}

func maxSubArray(a []int) int {
	tmp := 0
	max := -4096

	if len(a) == 1 {
		return a[0]
	}

	for i := 0; i < len(a); i++ {
		if tmp <= 0 {
			tmp = a[i]
		} else {
			tmp += a[i]
		}
		if tmp > max {
			max = tmp
		}
	}
	return max
}

/**
解题思路：
贪心的做法，可以线性时间解决。
如果之前的连续状态是正数，加上后面的值肯定更大，如果是负数就不要加了
即：
	f[i] = a[i] i==0, f[i-1] < 0
	f[i] = f[i-1] + a[i]

testcase:
9
-2 1 -3 4 -1 2 1 -5 4

out: 6
------------
2
-2, -1
-1
*/
