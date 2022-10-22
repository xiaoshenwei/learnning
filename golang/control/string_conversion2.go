package main

import (
	"fmt"
	"strconv"
)

func main() {
	orig := "123Q"
	// var newS string

	fmt.Println("The size ints is: %d", strconv.IntSize)

	if an, err := strconv.Atoi(orig); err != nil {
		fmt.Println("orig is not an int", an)
		return
	}
	// fmt.Println(an)
	// an = an + 5
	// newS = strconv.Itoa(an)
	// fmt.Println(newS)
}
