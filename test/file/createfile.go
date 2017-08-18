package main

import (
	"os"
	"fmt"
)

var (
	newFile *os.File
	err error
)

func main(){
	newFile, err = os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(newFile)
	newFile.Close()

	// 得到文件信息
	fileInfo, err := os.Stat("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.Size())
	fmt.Println(fileInfo.IsDir())
	fmt.Println(fileInfo.Mode())
	fmt.Println(fileInfo.Sys())
}
