package main

import (
	"fmt"

	"github.com/bkohler93/advent-of-code/2023/day-six/loader"
	"github.com/bkohler93/advent-of-code/2023/day-six/race"
)

func main() {
	input := loader.GetInput("input.txt",
		"2023", "6",
		"53616c7465645f5fb169fa6c83641e8c3c98857c5d60a9e9b0ae64ebeeb1ba35d5af9d61d1ad955673a0d1d4291bdcd4a216a15bcb464fbbb7cc41b8c050ea3c")

	races := race.GetRacesPartTwo(input)

	final := race.PartOne(races)

	fmt.Printf("Answer is %d\n", final)
}
