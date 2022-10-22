package main

import "fmt"

func main() {
	num := 20
	for i := 0; i < num; i++ {
		MyPrint(i)
	}

}

func MyPrint(num int) {
	switch {
	case (num%3 == 0 && num%5 == 0):
		fmt.Println("FizzBuzz")
	case num%3 == 0:
		fmt.Println("Fizz")
	case num%5 == 0:
		fmt.Println("Buzz")
	default:
		fmt.Println(num)
	}
}
