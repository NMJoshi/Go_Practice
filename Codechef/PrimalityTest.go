package main

import (
	"fmt"
	//"strings"
)

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)
		x := factor(n)
		if x == 2 {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}
func factor(n int) int {
	count := 1
	if n != 1 {
		for i := 2; i <= n; i++ {
			if n%i == 0 {
				count++
			}
		}
	}
	return count
}
