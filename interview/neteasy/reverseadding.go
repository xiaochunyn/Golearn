package main

import "fmt"
/**
*  网易2018年秋招笔试题 - 相反数
*  大致题意：
*  输入一个整数n(n >= 1 && n <= 10^5), 将n反转，求反转后的数和n的和
*  注：100反转之后为1
*
*  input:
*  203  （反转之后为302）
*
*  output：
*  505
*
*  分析：
*  主要是求N的反转数
*/
func main(){
	for {
		var val int
		_, err := fmt.Scanln(&val)
		if err != nil {
			fmt.Println(err)
		}
		reverse := 0
		left := val
		for left > 0{
			reverse = reverse * 10 + left % 10
			left = left / 10
		}
		fmt.Println(reverse)
		fmt.Println(val + reverse)
	}
}
