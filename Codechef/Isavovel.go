package main

import "fmt"

func main() {
	var c char
	fmt.Scan(&c)
	switch c {
	case A:
		fmt.Println("Vowel")
	case E:
		fmt.Println("Vowel")
	case I:
		fmt.Println("Vowel")
	case O:
		fmt.Println("Vowel")
	case U:
		fmt.Println("Vowel")
	default:
		fmt.Println("Constant")
	}
}
