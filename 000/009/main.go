package main

import (
	"errors"
	"fmt"
)

// There exists exactly one Pythagorean triplet for which a + b + c = 1000.
// Find the product abc.

const targetSum = 1000

// isTriplet returns whether the provided integers form a pythagorean triplet.
func isTriplet(a, b, c int) bool {
	return a*a+b*b == c*c
}

// findTriplet returns a pythagorean triplet that sums to sum, and an error if none exists.
func findTriplet(sum int) (a, b, c int, err error) {
	for c := sum - 2; c > 0; c-- {
		for b := sum - c - 1; b >= 1; b-- {
			a := sum - c - b
			if isTriplet(a, b, c) {
				return a, b, c, nil
			}
		}
	}
	err = errors.New("no triplet exists with that sum")
	return 0, 0, 0, err
}

func main() {
	a, b, c, err := findTriplet(targetSum)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(a * b * c)
}
