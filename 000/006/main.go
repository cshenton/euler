package main

import "fmt"

// Find the difference between:
// The sum of the squares of the first one hundred natural numbers and;
// The square of the sum.

const maxInt = 100

// Returns the sum of the squares of the numbers from 1 to max.
func sumSquare(max int) (s int) {
	for i := 1; i <= max; i++ {
		s += i * i
	}
	return s
}

// Returns the square of the sum of the numbers from 1 to max.
func squareSum(max int) (s int) {
	for i := 1; i <= max; i++ {
		s += i
	}
	return s * s
}

func main() {
	diff := squareSum(maxInt) - sumSquare(maxInt)
	fmt.Println(diff)
}
