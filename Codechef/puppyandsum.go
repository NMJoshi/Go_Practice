package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var d, n, sum int
		fmt.Scan(&d, &n)
		for i := 0; i < d; i++ {
			for j := 1; j <= n; j++ {
				sum += j
			}
			n = sum
			sum = 0
		}
		fmt.Println(n)
	}
}
