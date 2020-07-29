package main

import "fmt"

func main()  {
	fmt.Println("How are you doing?") 

	var response string 
	fmt.Scan(&response)

	fmt.Printf("I'm %v.", response) 

	fmt.Println("How are you doing?") 

	var response1 string 
	var response2 string 
	fmt.Scan(&response1)
	fmt.Scan(&response2)

	fmt.Printf("I'm %v %v", response1, response2) 

	fmt.Println("What would you like for lunch?")

	var food string
	fmt.Scan(&food)

	fmt.Printf("Sure, we can have %v for lunch.", food)
}