package main

import (
	"fmt"
)

func main() {
	var t, a, b, c, sum int
	//fmt.Print("Enter Test Cases:")
	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		//fmt.Println(i)
		fmt.Scan(&a, &b, &c)
		//fmt.Scanf("%d", &b)
		//fmt.Scanf("%d", &c)
		sum = a + b + c
		//fmt.Println(a, b, c, sum)
		if sum == 180 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

	}
}

//DONE
