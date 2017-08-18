package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func main(){
	file, err := os.OpenFile("test.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	//write to file
	txt := []byte("Bytes! \n")
	w, err := file.Write(txt)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Wrote %d bytes. \n", w)

	//-----------------ioutil.WriteFile--------------------
	err = ioutil.WriteFile("test1.txt", txt, 0666)
	if err != nil {
		fmt.Println(err)
	}

	dir, err := filepath.Abs(os.Args[0])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dir)

}
