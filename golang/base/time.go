package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)

	tn := time.Now().UTC()
	fmt.Println(tn)
	// 一周后时间
	week := 60 * 60 * 24 * 7 * 1e9
	week_from_now := t.Add(time.Duration(week))
	fmt.Println(week_from_now)
}
