package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "2017-08-30T15:03:12.102904695Z container start 044c3a99bd598c2efc4511ff0df7075e38dd5f5d451cfb5eb024f1e7ff707718 (image=ubuntu:uts, name=pod1"
	a := strings.Split(s, " container ")
	fmt.Println(a[0])
	fmt.Println(a[1])

	fmt.Println(strings.Split(a[1], " ")[1])
}
