package main

import "fmt"

// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

// i.e. for i in 1 to 20, x % i == 0

const stop = 1e10

func main() {
	for i := 1; i < stop; i++ {
		for j := 1; j <= 20; j++ {
			if i%j != 0 {
				break
			}
			if j == 20 {
				fmt.Println(i)
				return
			}
		}
	}
	// Better way, find mutual prime factors of the {1, ..., 20}, take product
}
