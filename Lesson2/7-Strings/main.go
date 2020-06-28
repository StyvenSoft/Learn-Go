package main

import "fmt"

func main() {
	
	var nameOfSong string
	var nameOfArtist string

	nameOfSong = "Stop Stop"
	nameOfArtist = "The Julie Ruin"

	var description string
	description = nameOfSong + " is by: " + nameOfArtist + "."
	fmt.Println(description)
	// Prints "Stop Stop is by: The Julie Ruin.


	// Define a string variable
	// called favoriteSnack
	
	// Assign a value to
	// favoriteSnack
	
	// Print out the message
	// "My favorite snack is "
	// followed by the value in
	// favoriteSnack
	var favoriteSnack string
	favoriteSnack = "delicious green broccoli"
	fmt.Println("My favorite snack is " + favoriteSnack)
}