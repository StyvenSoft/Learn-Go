package main

import "fmt"

func main()  {

	bored := true
	fmt.Println(!bored) // Prints false

	tired := false;
	fmt.Println(!tired) // Prints true

	readyToGo := true

	if !readyToGo {
		fmt.Println("Start the car!")
	} else {
		fmt.Println("What are we waiting for??")
	}
}