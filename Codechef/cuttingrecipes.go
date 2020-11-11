//UNCOMPELET.....
package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n, a, b int
		fmt.Scan(&n)
		item := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&item[i])
		}
		for i := 0; i < n; i++ {
			var j int
			for j = 2; j < n+1; j++ {
				if item[i]%j == 0 {
					flag = 1
					break
				} else {

				}

			}
		}
	}
}
