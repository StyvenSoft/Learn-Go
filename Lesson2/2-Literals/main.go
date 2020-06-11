// We can perform arithmetic in Go with literals (or named values, covered in the next exercise) using 
// the following operators:

/*
+ to add
- to subtract
* to multiply
/ to divide
% to take the remainder (the modulus operator) between two numbers.
*/

package main

import "fmt"

func main() {
  // Add a fmt.Println() statement
  // that prints 2235 * 1231
  fmt.Println(2235 * 1231)
  fmt.Println(20 * 3) // Prints: 60
  fmt.Println(55.21 / 5) // Prints: 11.042
  fmt.Println(9 % 2) // Prints: 1
}