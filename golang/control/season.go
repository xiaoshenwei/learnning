package main

import "fmt"

func main() {
	month := 51
	switch month {
	case 1, 2, 3:
		fmt.Println("春季")
	case 4, 5, 6:
		fmt.Println("夏季")
	case 7, 8, 9:
		fmt.Println("球季")
	case 10, 11, 12:
		fmt.Println("冬季")
	default:
		fmt.Println("非法输入")
	}
}
