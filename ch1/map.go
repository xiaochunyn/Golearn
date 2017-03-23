package main

import "fmt"

func main() {
	m := map[string]int{
		"a": 1, "b": 2, "c": 3,
	}

	for k, v := range m {
		fmt.Printf("Key: %s, value: %d \n", k, v)
	}

	fmt.Println(m["a"])
	m["d"] = 4
	m["e"] = 5

	fmt.Println(m["e"])

	delete(m, "e")
	fmt.Println(m)

	v, ok := m["f"]
	fmt.Println(v)
	fmt.Println(ok)

}
