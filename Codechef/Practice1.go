/*Consider a currency system in which there are notes of six denominations, namely,
Rs. 1, Rs. 2, Rs. 5, Rs. 10, Rs. 50, Rs. 100. If the sum of Rs. N is input,
write a program to computer smallest number of notes that will combine to give Rs. N.*/
package main

import "fmt"

var t, n int

func main() {

	//fmt.Println("Enter value of t")
	fmt.Scanf("%d", &t)
	fmt.Print("A")

	for i := 0; i < t; i++ {
		//fmt.Println("enter value of n")
		fmt.Print("B")
		fmt.Scanf("%d", &n)
		fmt.Print("C")
		count := 0
		fmt.Print("D")
		for i := 0; i <= n; i++ {
			if n >= 100 {
				count += n / 100
				n = n % 100
			} else if n >= 50 {
				count += n / 50
				n = n % 50
			} else if n >= 10 {
				count += n / 10
				n = n % 10
			} else if n >= 5 {
				count += n / 5
				n = n % 5
			} else if n >= 2 {
				count += n / 2
				n = n % 2
			} else if n >= 1 {
				count += n
			} else {
				fmt.Print()
			}
		}
		fmt.Print("E")

		fmt.Print(count)
		fmt.Print("F")

	}
}
