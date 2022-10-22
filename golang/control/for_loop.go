package main

import (
	"fmt"
)

func main() {
	// for i := 0; i < 15; i++ {
	// 	fmt.Println(i)
	// }
	for i := 1; i < 25; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("G")
		}
		fmt.Println()
	}
	// for i := 1; i < 25; i++ {
	// 	fmt.Println(strings.Repeat("G", i))
	// }
}
