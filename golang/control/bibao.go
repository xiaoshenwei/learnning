package main

import (
	"fmt"
	"strings"
)

func main() {
	shF := MakeAddSuffix(".sh")
	bashF := MakeAddSuffix(".bash")

	fmt.Println(shF("aaa"))
	fmt.Println(bashF("aaa"))
}

func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
