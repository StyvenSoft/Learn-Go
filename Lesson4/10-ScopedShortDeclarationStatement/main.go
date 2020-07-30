package main

import "fmt"

func main() {

	x := 8
	y := 9
	if product := x * y; product > 60 {
	fmt.Println(product, "  is greater than 60")
	}

	switch season := "summer" ; season {
	case "summer":
		fmt.Println("Go out and enjoy the sun!")
	}

	success := true

	if success {
		fmt.Println("We're rich!")
	} else {
		fmt.Println("Where did we go wrong?")
	}

	amountStolen := 50000
	numOfThieves := 5

	switch numOfThieves {
	case 1:
		fmt.Println("I'll take all $", amountStolen)
	case 2:
	  fmt.Println("Everyone gets $", amountStolen/2)
	case 3:
		fmt.Println("Everyone gets $", amountStolen/3)
	case 4:
	  fmt.Println("Everyone gets $", amountStolen/4)
	case 5:
		fmt.Println("Everyone gets $", amountStolen/5)
	default:
		fmt.Println("There's not enough to go around...")
	}
}