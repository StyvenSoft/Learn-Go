package main

import "fmt"

func addHundred (numPtr *int) {
	*numPtr += 100
}

// Change brainwash to have a pointer parameter
func brainwash(saying *string) {
	// Dereference saying below: 
	*saying = "Beep Boop."
}

func main() {

	x := 1
	addHundred(&x)
	fmt.Println(x) // Prints 101

	greeting := "Hello there!"
	
	// Call your brainwash() below:
	brainwash(&greeting)
	
	fmt.Println("greeting is now:", greeting)
}