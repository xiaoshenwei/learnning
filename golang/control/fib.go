package main

import "fmt"

func main() {
	// for i := 0; i < 20; i++ {
	// 	// res := Fib(i)
	// 	// fmt.Println(res)
	// 	res, idx := FibIdx(i)
	// 	fmt.Printf("Num: %d, Idx: %d\n", res, idx)
	// }
	MyPrint(10)
}

func Fib(num int) int {
	if num <= 1 {
		return 1
	}
	return Fib(num-1) + Fib(num-2)
}

func FibIdx(num int) (res int, idx int) {
	if num <= 1 {
		return 1, num
	}
	res1, _ := FibIdx(num - 1)
	res2, _ := FibIdx(num - 2)

	return res1 + res2, num
}

func MyPrint(num int) {
	if num == 1 {
		return
	}
	fmt.Println(num)
	MyPrint(num - 1)
}
