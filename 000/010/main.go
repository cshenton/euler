package main

import (
	"fmt"
	"math"
)

// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
// Find the sum of all the primes below two million.

// Let's make a prime sieve.
// We want to represent the set of all numbers not divisible by any previous number
// So lets build a cascading sequence of channels that check for the presence of the nth prime.

const (
	maxPrime = 2e6
)

// Let's not use the intentionally bad channel sieve.

// checkPrimes returns the non-prime statuses of all numbers up to max.
func checkPrimes(max int) (s []bool) {
	s = make([]bool, max+1)
	s[0] = true
	s[1] = true
	s[2] = false

	for i := 4; i <= max; i += 2 {
		s[i] = true
	}

	limit := int(math.Sqrt(float64(max))) + 1
	for i := 3; i < limit; i += 2 {
		if !s[i] {
			for j := i * i; j < max; j += i {
				s[j] = true
			}
		}
	}
	return s
}

// sumPrimes returns the sum of all primes less than max
func sumPrimes(max int) (sum int) {
	s := checkPrimes(max)
	for i := 0; i <= max; i++ {
		if !s[i] {
			sum += i
		}
	}
	return sum
}

func main() {
	sum := sumPrimes(maxPrime)
	fmt.Println(sum)
}
