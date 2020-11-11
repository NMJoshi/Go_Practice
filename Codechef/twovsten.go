//SUCCESSFULL
package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)
		unit := n % 10
		if unit == 0 {
			fmt.Println("0")
		} else if unit == 5 {
			fmt.Println("1")
		} else {
			fmt.Println("-1")
		}
	}
}
