package main

import (
	"fmt"

	"github.com/csharpknife2002/golang/my_hello/morestrings"
	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Printf("hello, world from me\n")
	fmt.Printf(morestrings.ReverseString("I am Chen LIN"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
