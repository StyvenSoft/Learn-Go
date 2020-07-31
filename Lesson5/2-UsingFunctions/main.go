package main

import "fmt"

func summonNicole() {
  fmt.Println("Hey Nicole, get over here!")
}

// Define eatTacos() here:
func eatTacos()  {
	fmt.Println("Yum!")
}

func main() {
  // We call our function for the first time
  summonNicole() 

  eatTacos()
}