package main

import "fmt"

func main()  {
	clothingChoice := "sweater"

	if clothingChoice == "shirt" {
		fmt.Println("We have shirts in S and M only.")
	} else if clothingChoice == "polos" {
		fmt.Println("We have polos in M, L, and XL.")
	} else if clothingChoice == "sweater" {
		fmt.Println("We have sweaters in S, M, L, and XL.")
	} else {
		fmt.Println("Sorry, we don't carry that.")
	}

	name := "H. J. Simp"

	switch name {
	case "Butch":
		fmt.Println("Head to Robbers Roost.")
	case "Bonnie":
		fmt.Println("Stay put in Joplin.")
	default:
		fmt.Println("Just hide!")
	}
	
}