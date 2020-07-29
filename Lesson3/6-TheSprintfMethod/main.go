package main

import "fmt"

func main()  {
	correctAns := "A"
	answer := fmt.Sprintf("And the correct answer is… %v!", correctAns)

	fmt.Print(answer) // Prints: And the correct answer is… A!

	template := "I wish I hab a %v."
	pet := "puppy"

	var wish string

	wish = fmt.Sprintf(template, pet)
  
  	fmt.Println(wish)
}

// type in go run main.go and press enter.