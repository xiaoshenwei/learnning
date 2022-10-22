package main

import "fmt"

func main() {
	var arr1 [5]int
	// var arr2 [5]int
	fmt.Printf("addr: %p\n", &arr1)
	Pp(arr1)
	// Pp(arr2)
}

func Pp(arr [5]int) {
	fmt.Printf("Ppaddr: %p\n", &arr)
}
