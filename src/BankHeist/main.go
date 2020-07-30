package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	rand.Seed(time.Now().UnixNano())
	isHeistOn := true
	eludedGuards := rand.Intn(100)

	if eludedGuards >= 50 {
		fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
	} else {
		isHeistOn = false
		fmt.Println("Plan a better disguise next time?")
	} 

	openedValue := rand.Intn(100)

	if isHeistOn && openedValue >= 70 {
		fmt.Println("Grab and GO!")
	} else if isHeistOn {
		isHeistOn = false
		fmt.Println("Did they superglue the vault?")
	}

	leftSafely := rand.Intn(5)

	if isHeistOn {
		switch leftSafely {
		case 0:
			isHeistOn = false
			fmt.Println("Looks like you tripped an alarm!")
		case 1:
			isHeistOn = false
			fmt.Println("Turns out vault doors don't open from the inside...")
		case 2:
			isHeistOn = false
			fmt.Println("Right, banks have security cameras...")
		case 3:
			isHeistOn = false
			fmt.Println("It's just bab luck that cops also need to use banks.")
		default:
			fmt.Println("Start the getaway car!")
		}
	}

	if isHeistOn {
		amtStolen := 10000 + rand.Intn(1000000)
		fmt.Println("$", amtStolen, "not bab!")
	}

	fmt.Println(isHeistOn)
}