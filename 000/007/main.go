package main

import "fmt"

// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
// What is the 10 001st prime number?

// Let's make a prime sieve.
// We want to represent the set of all numbers not divisible by any previous number
// So lets build a cascading sequence of channels that check for the presence of the nth prime.

const (
	primeNum = 10001
	workers  = 10
)

// Sends the sequence 2,3,4... to ch.
func gen(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

// filter sends all values from in to out not divisible by prime.
func filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

// nPrime returns the nth prime number.
func nPrime(n int) (p int) {
	ch := make(chan int)
	go gen(ch)
	for i := 0; i < n; i++ {
		p = <-ch
		nch := make(chan int) // channel of everything not divisible up to prime
		go filter(ch, nch, p)
		ch = nch
	}
	return p
}

func main() {
	prime := nPrime(primeNum)
	fmt.Println(prime)
}
