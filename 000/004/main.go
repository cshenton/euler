package main

import (
	"fmt"
	"strconv"
)

// A palindromic number reads the same both ways.
// The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
// Find the largest palindrome made from the product of two 3-digit numbers.

// reverse reverses the input string.
func reverse(s string) string {
	if s != "" {
		return reverse(s[1:]) + s[:1]
	}
	return s

}

// isPalindrome return whether the input integer is a palindrome.
func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	return reverse(s) == s
}

func main() {
	// Iterate along size ordered diagonals of x, y pairs.
	for i := 0; i <= 889; i++ {
		for j := 0; j <= i+1; j++ {
			x := 999 - i + j
			y := 999 - j
			if isPalindrome(x * y) {
				fmt.Println(x, y, x*y)
				return
			}
		}
	}

	fmt.Println("no palindromes found")
	return
}
