package main

import (
	"fmt"
	"reflect"
)

type AA struct {
	s string
}
// 深度等于判断案例
func main() {

	a1 := AA{s: "abc"}
	a2 := AA{s: "abc"}
	if reflect.DeepEqual(a1, a2) {
		fmt.Println(a1, "==", a2)
	}

	b1 := []int{1, 2}
	b2 := []int{1, 2}
	if reflect.DeepEqual(b1, b2) {
		fmt.Println(b1, "==", b2)
	}

	c1 := map[string]int{"a": 1, "b": 2}
	c2 := map[string]int{"a": 1, "b": 2}
	if reflect.DeepEqual(c1, c2) {
		fmt.Println(c1, "==", c2)
	}
}
