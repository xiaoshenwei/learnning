package main

import "fmt"

func main() {
	i1 := 5
	fmt.Printf("Int: %d, Id: %p", i1, &i1)

	intp := &i1
	fmt.Printf("v of Id: %d\n", *intp)
}
