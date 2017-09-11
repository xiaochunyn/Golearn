package main

import (
	"fmt"
)

/**
时间限制：1秒
空间限制：32768K
给定一个十进制数M，以及需要转换的进制数N。将十进制数M转化为N进制数
输入描述:
输入为一行，M(32位整数)、N(2 ≤ N ≤ 16)，以空格隔开。
输出描述:
为每个测试实例输出转换后的数，每个输出占一行。如果N大于9，则对应的数字规则参考16进制（比如，10用A表示，等等）
输入例子1:
7 2
输出例子1:
111
解析：

（1）注意负数，需要转换成正数；
（2）如果是-2147483648，转换位正数是214748648，超出了32位int范围，得用unsigned int或者long 。
*/
func main() {
	for {
		var m, n int
		fmt.Scanf("%d %d", &m, &n)
		fmt.Printf("%d %d \n", m, n)
		if n < 2 || n > 16 {
			fmt.Println("error")
			continue
		}
		decToN(m, n)
	}
}

func decToN(m, n int) {
	factor := "0123456789ABCDEF"

	if m == 0 {
		fmt.Println(0)
		return
	}

	// 如果m为正数，flag为true
	flag := true
	var val int

	if m < 0 {
		flag = true
		val = ^m + 1
	} else {
		val = m
	}

	index := 31
	str := [32]byte{0}

	for val != 0 {
		str[index] = factor[val%n]
		index--
		val /= n
	}
	if !flag {
		fmt.Print("-")
	}
	for i := index + 1; i < 32; i++ {
		fmt.Printf("%q", str[i])
	}
	fmt.Println()
}

/**
#include <iostream>
using namespace std;

string factor = "0123456789ABCDEF";

void decToN(int num,int n)
{
    if(num==0){
        cout<<0;
        return;
    }
    bool flag = true;
    unsigned int m = 0;
    if(num<0){
        m = ~num + 1;   //注意-2147483648 ~ 2147483647
        flag = false;
    }
    else
        m = num;
    int index = 31;
    char str[32]={0};
    while(m!=0){
        str[index--] = factor[m%n];
        m /= n;
    }
    if(!flag)
        cout<<'-';
    for(int i=index+1; i<32; i++)
        cout<<str[i];
}

int main()
{
    int m,n;
    cin>>m>>n;
    if(n<2 || n>16)
        return 0;
    decToN(m,n);
    return 0;
}

*/
