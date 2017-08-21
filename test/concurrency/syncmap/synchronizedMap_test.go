package main

import (
	"fmt"
	"testing"
)

func Test_SynchonizedMap(t *testing.T) {
	m := NewSynchronizedMap()
	m.Put("Tom", 19)
	m.Put("Luxy", 18)
	m.Put("Dillon", 17)

	tom := m.Get("Tom")
	if tom != 19 {
		t.Errorf("Failed to get Tom's age. ")
	}

	m.Delete("Tom")
	if ok := m.Get("Tom"); ok != nil {
		t.Errorf("Failed to delete Tom.")
	}

	m.Each(func(k interface{}, v interface{}) {
		v = v.(int) + 1
		fmt.Println(k, v)
	})

}
