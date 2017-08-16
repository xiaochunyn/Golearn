package main

import "fmt"

func main(){
	var name string
	n, err := fmt.Scanf("%s", &name)
	fmt.Println(n, err, name)
}