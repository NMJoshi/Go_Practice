package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n, temp, reverse, reminder int
		fmt.Scan(&n)
		temp = n
		for i := n; i > 0; i++ {
			reminder = n % 10
			reverse = reverse*10 + reminder
			n = n / 10
			if n == 0 {
				break
			}
		}
		if temp == reverse {
			fmt.Println("wins")
		} else {
			fmt.Println("loses")
		}
	}
}
