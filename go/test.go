package main

import "fmt"

type T struct {
	a string
	b int32
}

func main() {

	s := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(s[2:5])

}

func tt(a *int32) func() int32 {

	f := func() int32 {
		return *a * 2
	}
	return f
}
