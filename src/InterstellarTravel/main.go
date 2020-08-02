package main

import "fmt"

// Create the function fuelGauge() here

func fuelGauge (fuel int)  {
	fmt.Println("You have", fuel, "gallons of fuel left!")
}

// Create the function calculateFuel() here

func calculateFuel(planet string) int {
	var fuel int
	switch planet {
	case "Venus":
		fuel = 300000
	case "Mercury":
		fuel = 500000
	case "Mars":
		fuel = 700000
	default:
		fuel = 0
	}
	return fuel
}

// Create the function greetPlanet() here

func greetPlanet(planet string) {
	fmt.Println("Welcome to planet", planet)
}

// Create the function cantFly() here
func cantFly ()  {
	fmt.Println("We do not have the avaible fuel to fly there")
}

// Create the function flyToPlanet() here
func flyToPlanet(planet string, fuel int) int {
	var fuelRemaining, fuelConst int
	fuelRemaining = fuel
	fuelConst = calculateFuel(planet)

	if fuelRemaining >= fuelConst {
		greetPlanet(planet)
		fuelRemaining -= fuelConst
	} else {
		cantFly()
	}
	return fuelRemaining
}

func main() {
	// Test your functions!
	//fuelGauge(7000)
	// Create `planetChoice` and `fuel`
	//fmt.Println(calculateFuel("Mars"))
	//greetPlanet("Venus")
	//cantFly()

	var fuel int
	fuel = 1000000

	var planetChoice string
	planetChoice = "Venus"
	// And then liftoff!
	
	fuel = flyToPlanet(planetChoice, fuel)

	fuelGauge(fuel)
  }