package main

import (
	"fmt"
	"time"
)

type T struct {
	a string
	b int32
}

func main() {

	s := time.Unix(1632257515, 0).Format("2006-01-02 15:04:05")
	fmt.Println(s)

}

func tt(a *int32) func() int32 {

	f := func() int32 {
		return *a * 2
	}
	return f
}
