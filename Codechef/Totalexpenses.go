package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var a, b, sum float64
		fmt.Scan(&a, &b)
		if a > 1000 {
			sum = a * b
			temp := sum / 10
			sum = sum - temp
		} else {
			sum = a * b
		}
		fmt.Println(sum)
	}
}
