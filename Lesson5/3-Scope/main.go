package main

import "fmt"

func performAddition() {
	x := 5
	y := 7
	fmt.Println("The sum of", x, "and", y, "is", x + y)
	fmt.Println("What if", x, "was different?")
}

func startGame() {
	instructions := "Press enter to start..." 
	
	fmt.Println(instructions)
}
  
func main() {
	performAddition()
	startGame()
}