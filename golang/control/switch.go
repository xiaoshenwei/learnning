package main

import "fmt"

func main() {
	i := 0
	switch i {
	case 0:
		fallthrough
	case 1:
		fmt.Println(i)
	case 2:
		fmt.Println(i + 1)
	}
	k := 6
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
		fallthrough
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
		// fallthrough
	default:
		fmt.Println("default case")
	}
}
