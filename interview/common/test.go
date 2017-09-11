package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	s := "/proc/26380/ns/net"
	a := strings.Split(s, "/")
	fmt.Println(len(a), a[2])
	fmt.Println(time.Friday)
	str := [32]byte{0}
	fmt.Println(str[3])
	//
	//s := "asdf"
	//fmt.Println([]rune(s))
	//b := []byte{'a', 's', 'd'}
	//fmt.Printf("% q", b)
	//fmt.Printf("%q", s[1])
}
