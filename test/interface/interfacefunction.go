package main

import "fmt"

// Handler define a interface first
type Handler interface {
	Do(k, v interface{})
}

// HandleFunc define a function type
type HandleFunc func(k, v interface{})

// Do HandleFunc to imply the Handle interface
func (f HandleFunc) Do(k, v interface{}) {
	f(k, v)
}

// Each traverse the map
func Each(m map[interface{}]interface{}, h Handler) {
	if m != nil && len(m) > 0 {
		for k, v := range m {
			h.Do(k, v)
		}
	}
}

// EachFunc
func EachFunc(m map[interface{}]interface{}, f func(k, v interface{})) {
	Each(m, HandleFunc(f))
}

func selfInfo(k, v interface{}) {
	fmt.Printf("Hello, my name is %s,\t I'm %d years old. \n", k, v)
}

func main() {
	student := make(map[interface{}]interface{})
	student["Tom"] = 10
	student["Lucy"] = 20
	student["Lily"] = 18

	EachFunc(student, selfInfo)
}
