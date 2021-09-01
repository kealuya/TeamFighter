package main

import "fmt"

type T struct {
	a string
	b int32
}

func main() {

	var a = &T{
		a: "aaa",
		b: 111,
	}
	f := tt(&a.b)
	a.b = 444
	fmt.Println(f())
}

func tt(a *int32) func() int32 {

	f := func() int32 {
		return *a * 2
	}
	return f
}
