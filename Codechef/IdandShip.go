package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var ch string
		fmt.Scan(&ch)
		switch ch {
		case "B", "b":
			fmt.Println("BattleShip")
		case "C", "c":
			fmt.Println("Cruiser")
		case "D", "d":
			fmt.Println("Destroyer")
		case "F", "f":
			fmt.Println("Frigate")
		default:
			fmt.Println("")
		}
	}
}
