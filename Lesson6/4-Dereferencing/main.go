package main

import "fmt"

func main() {

	lyrics := "Moments so dear" 
	pointerForStr := &lyrics

	*pointerForStr = "Journeys to plan" 

	fmt.Println(lyrics) // Prints: Journeys to plan

	star := "Polaris"
	
	starAddress := &star
	
	// Add your code below:
	*starAddress = "Sirius"
	
  
  fmt.Println("The actual value of star is", star)
}