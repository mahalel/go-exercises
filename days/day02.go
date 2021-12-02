package days

import (
	"fmt"
	"go/exercises/inputs"
	"strconv"
	"strings"
)

// https://adventofcode.com/2021/day/2
// Two : advent of code, day one part1 and 2

// This function moves the submarine
func MoveSub(direction string, distance int, aim int) (int, int, int) {
	var horizontal int = 0
	var depth int = 0

	if direction == "down" {
		aim += distance
	} else if direction == "up" {
		aim -= distance
	} else if direction == "forward" {
		horizontal += distance
		depth += aim * distance
	}
	return horizontal, depth, aim
}

func Two() {
	// Load our input data
	var inputSlice []string = inputs.Day2
	var horizontal_total, depth_total, aim_total int = 0, 0, 0

	for i := range inputSlice {
		// Split our input lines into an array with two elements
		v := strings.Split(inputSlice[i], " ")
		// Convert the second value of the array to int
		y, err := strconv.Atoi(v[1])
		if err != nil {
		}
		// Move the submarine and assign the returned values to new variables
		var horizontal_val, depth_val, aim_val = MoveSub(v[0], y, aim_total)
		// Re-assign the values or increment
		horizontal_total += horizontal_val
		depth_total += depth_val
		aim_total = aim_val
	}
	fmt.Println(horizontal_total * depth_total)
}
