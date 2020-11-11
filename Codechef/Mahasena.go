package main

import "fmt"

func main() {
	var n, even, odd int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	for i := 0; i < n; i++ {
		if a[i]%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	if even > odd {
		fmt.Print("READY FOR BATTLE")
	} else {
		fmt.Print("NOT READY")
	}
}
