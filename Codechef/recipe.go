//SUCCESSFULL
package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n, gcd int
		fmt.Scan(&n)
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&a[i])
		}
		//fmt.Print(a)
		//common divisor
		gcd = find_gcd(a[0], a[1])
		for i := 2; i < n; i++ {
			gcd = find_gcd(gcd, a[i])
		}
		for i := 0; i < n; i++ {
			fmt.Print(a[i]/gcd, " ")
		}
		fmt.Println()
	}
}
func find_gcd(x, y int) int {
	if x == 0 {
		return y
	} else {
		return find_gcd(y%x, x)
	}
}
