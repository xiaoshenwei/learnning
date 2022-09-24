package main

import "fmt"

func main() {
	s := "Good Bye!!!"
	p := &s
	*p = "Meet again"
	fmt.Printf("P %p\n", p)
	fmt.Printf("s %s\n", s)
	fmt.Printf("*p %s\n", *p)
}
