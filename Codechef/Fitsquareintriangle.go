package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var b int
		fmt.Scan(&b)
		b = b - 2
		b = b / 2
		ans := (b * (b + 1)) / 2
		fmt.Println(ans)
	}
}
