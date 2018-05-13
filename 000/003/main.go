package main

import (
	"fmt"
	"math"
)

// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ?

const target = 600851475143

// factors returns the ordered prime factors of input integer x.
func factors(x int) (f []int) {
	f = make([]int, 0)

	// Reduce x to an odd number
	for x%2 == 0 {
		f = append(f, 2)
		x = x / 2
	}

	// Find all remaining factors
	for i := 3; float64(i) <= math.Sqrt(float64(x)); i += 2 {
		for x%i == 0 {
			f = append(f, i)
			x = x / i
		}
	}

	// Final check if remaining x is prime
	if x > 2 {
		f = append(f, x)
	}

	return f
}

func main() {
	f := factors(target)
	fmt.Println(f[len(f)-1])
}
