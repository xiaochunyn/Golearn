package main

import "fmt"

func main() {
	m := make(map[int]string)
	m[1] = "tom"
	m[2] = "cat"
	fmt.Println(m[1])
	fmt.Println(m[2])

	m1 := m
	fmt.Printf("%x \n", &m)
	fmt.Printf("%x \n", &m1)
	m1[1] = "Lily"
	fmt.Println(m[1])
	fmt.Println(m[2])
	fmt.Println(m1[1])
	fmt.Println(m1[2])
}

/**
map 是地址引用，如下

tom
cat
&map[1:746f6d 2:636174]
&map[1:746f6d 2:636174]
Lily
cat
Lily
cat
*/
