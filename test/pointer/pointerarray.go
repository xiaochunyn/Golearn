package main

import "fmt"

func main(){
	a := []int{10,20,30}
	var ptr [3]*int

	for i, v := range a {
		//ptr[i] = &a[i]
		ptr[i] = &v 		// 切勿这样操作，因为v在循环结构中定义，&v是取地址操作
		fmt.Println(v)
	}

	for i, v := range ptr {
		fmt.Printf("a[%d] = %d \n", i, *v)
	}
}
