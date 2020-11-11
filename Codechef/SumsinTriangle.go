package main

import "fmt"

func main() {
	var n int
	fmt.Println("Test case:")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var r int
		fmt.Println("No of Rows:")
		fmt.Scan(&r)
		var a [100][100]int
		for i := 0; i < r; i++ {
			for j := 0; j <= i; j++ {
				fmt.Scan(&a[i][j])
			}
		}
		/*for i := 0; i < r; i++ {
			for j := 0; j <= i; j++ {
				add[i][j] = a[i][j]
			}
		}*/
		for i := r; i > 0; i-- {
			for j := 0; j <= i-1; j++ {
				if a[i][j] > a[i][j+1] {
					a[i-1][j] = a[i-1][j] + a[i][j]
				} else {
					a[i-1][j] = a[i-1][j] + a[i][j+1]
				}
			}
		}
		fmt.Println(a[0][0])
	}
}

/*func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}*/
/*func Print(){
	for j := 0; j < r; j++ {
		for k := 0; k <= j; k++ {
			fmt.Print(a[j][k], " ")
		}
		fmt.Print("\n")
	}
}*/
