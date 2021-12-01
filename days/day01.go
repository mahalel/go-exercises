package days

import (
	"fmt"
	"go/exercises/inputs"
)

// https://adventofcode.com/2021/day/1
// One : advent of code, day one part1 and 2

func SonarSweep(input []int) {
	var counter int = 0

	for i, value := range input {
		if i == len(input)-1 {
			break
		}
		if value < input[i+1] {
			counter++
		}
	}
	fmt.Println(counter)
}

func One() {
	inputSlice := inputs.Day1
	var newArr []int
	for i := range inputSlice {
		if i < len(inputSlice)-2 {
			x := inputSlice[i : i+3]
			slidingWindowSum := 0
			for _, v := range x {
				slidingWindowSum += v
			}
			newArr = append(newArr, slidingWindowSum)
		}
	}
	SonarSweep(inputSlice)
	SonarSweep(newArr)
}
