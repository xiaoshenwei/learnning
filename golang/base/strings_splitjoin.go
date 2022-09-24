// 拼接分割字符串
package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "The quick brown fox jumps over the lazy dog"
	sl := strings.Fields(str1)
	for i, v := range sl {
		fmt.Println(i, v)
	}

	sl1 := strings.Split(str1, " ")
	for _, v := range sl1 {
		fmt.Println(v)
	}

	str2 := strings.Join(sl1, "_")
	fmt.Println(str2)
}
