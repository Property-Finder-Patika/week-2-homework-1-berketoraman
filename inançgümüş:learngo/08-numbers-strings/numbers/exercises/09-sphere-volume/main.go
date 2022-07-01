package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

// EXERCISE: Sphere Volume
//  1. Get the radius from the command-line
//  2. Convert it to a float64
//  3. Calculate the volume of a sphere

// SPHERE VOLUME FORMULA
//  https://en.wikipedia.org/wiki/Sphere#Enclosed_volume

// RESTRICTION
//  Use `math.Pow` to calculate the volume

// EXPECTED OUTPUT
//  go run main.go 10
//    4188.79

//  go run main.go .5
//    0.52

func main() {
	var radius, vol float64

	radius, _ = strconv.ParseFloat(os.Args[1], 64)

	vol = (4 * math.Pi * math.Pow(radius, 3)) / 3

	fmt.Printf("radius: %g -> volume: %.2f\n", radius, vol)
}
