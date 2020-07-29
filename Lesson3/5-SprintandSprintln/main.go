package main

import "fmt"

func main()  {
	grade := "100"
	compliment := "Great job!"
	teacherSays := fmt.Sprint("You scored a ", grade, " on the test! ", compliment)

	fmt.Print(teacherSays)
	// Prints: You scored a 100 on the test! Great job!

	step1 := "Breathe in..."
	step2 := "Breathe out..."

	meditation := fmt.Sprintln(step1, step2)

	fmt.Printf(meditation)
}