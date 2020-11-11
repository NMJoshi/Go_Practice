//TIME EXCIDE
//REQUIRED 1.0 SEC
//HERE 1.01 SEC
package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		a1 := a
		b1 := b
		//fmt.Scan(&b)
		count := 2
		gcd := 1
		for j := 0; j < a || j < b; j++ {
			if a%count == 0 && b%count == 0 {
				a /= count
				b /= count
				gcd *= count
			} else {
				count++
			}
		}
		//gcd := gcd(a, b)
		lcm := (a1 * b1) / gcd
		fmt.Println(gcd, lcm)
	}
}

/*func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}*/
/*func gcd(a, b int) int {
	count := 2
	gcd := 1
	for j := 0; j < a || j < b; j++ {
		if a%count == 0 && b%count == 0 {
			a /= count
			b /= count
			gcd *= count
		} else {
			count++
		}
	}
	return gcd
}*/

/*func lcm(a, b int) int {
	lcm := 1
	for count := 2; count < max(a, b)+1; {
		if a%count == 0 && b%count == 0 {
			a /= count
			b /= count
			lcm *= count
		} else if a%count == 0 && a != 1 {
			a /= count
			lcm *= count
		} else if b%count == 0 && b != 1 {
			b /= count
			lcm *= count
		} else {
			count++
		}
	}

	return lcm
}*/
