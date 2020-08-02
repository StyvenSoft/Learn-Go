package main

import "fmt"

func main() {

	var pointerForInt *int

	minutes := 525600

	pointerForInt = &minutes

	fmt.Println(pointerForInt) // Prints 0xc000018038

	star := "Polaris"
	
	// Add your code below:
	
	starAddress := &star
  
	fmt.Println("The address of star is", starAddress)
}