package main

import (
	"fmt"
	"strconv"
)

func main() {
	num_str01 := "123a"
	num, ok := strconv.Atoi(num_str01)
	if ok != nil {
		fmt.Println("字符串转换失败")
	}
	fmt.Println(num, ok)

	num02 := 123
	num02_str := strconv.Itoa(num02)
	fmt.Println(num02_str)
}
