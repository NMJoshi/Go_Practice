package main

import (
  "fmt"
)

func main() {
  var t int
  fmt.Scan(&t)
  for i := 0; i < t; i++ {
    var a, b int
    fmt.Scan(&a, &b)
    gcd := gcd(a, b)
    lcm := (a * b) / gcd
    fmt.Println(gcd, lcm)
  }

  //fmt.Println(LCM(10, 15, 20))
  //fmt.Println(LCM(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}
func gcd(a, b int) int {
  for b != 0 {
    t := b
    b = a % b
    a = t
  }
  return a
}

/*func lcm(a, b int) int {
  result := a * b / GCD(a, b)

  for i := 0; i < len(integers); i++ {
    result = LCM(result, integers[i])
  }

  return result
}*/
