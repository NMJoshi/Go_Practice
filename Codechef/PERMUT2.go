//SUCCESSFULL
package main

import (
	"fmt"
	"reflect"
)

func main() {
	for i := 0; i < 1; {
		var n int
		fmt.Scan(&n)
		if n != 0 {
			a := make([]int, n+1)
			arr := make([]int, n+1)
			for i := 1; i <= n; i++ {
				fmt.Scan(&a[i])

			}
			for i := 1; i <= n; i++ {
				arr[a[i]] = i
			}
			if reflect.DeepEqual(a, arr) {
				fmt.Println("ambiguous")
			} else {
				fmt.Println("not ambiguous")
			}
		} else {
			i++
		}
	}
}
