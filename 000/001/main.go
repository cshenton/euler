package main

import "fmt"

// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9.
// The sum of these multiples is 23.

// Find the sum of all the multiples of 3 or 5 below 1000.

const max = 1000

var divisors = []int{3, 5}

func main() {
	multiples := []int{}
	result := 0

	for i := 1; i < max; i++ {
		for _, d := range divisors {
			if i%d == 0 {
				multiples = append(multiples, i)
				break
			}
		}
	}

	for _, m := range multiples {
		result += m
	}

	fmt.Println(result)
}
