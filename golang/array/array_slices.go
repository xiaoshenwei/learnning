package main

import "fmt"

func main() {
	var arr1 [5]int
	var slice1 []int = arr1[:]
	for _, v := range slice1 {
		fmt.Println(v)
	}
}
