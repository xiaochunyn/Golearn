package main

import "fmt"

/**
素数对
时间限制：1秒
空间限制：32768K
给定一个正整数，编写程序计算有多少对质数的和等于输入的这个正整数，并输出结果。输入值小于1000。
如，输入为10, 程序应该输出结果为2。（共有两对质数的和为10,分别为(5,5),(3,7)）
输入描述:
输入包括一个整数n,(3 ≤ n < 1000)

输出描述:
输出对数

输入例子1:
10

输出例子1:
2
*/
func main() {
	var val int
	var pair int
	fmt.Scanln(&val)
	for i := 0; i <= val/2; i++ {
		if isPrime(i) && isPrime(val-i) {
			pair++
		}
	}
	fmt.Println(pair)
}

func isPrime(v int) bool {
	for i := 2; i < v; i++ {
		if v%i == 0 {
			return false
		}
	}
	return true
}
