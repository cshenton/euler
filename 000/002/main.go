package main

import "fmt"

// Each new term in the Fibonacci sequence is generated by adding the previous two terms.
// By starting with 1 and 2, the first 10 terms will be:

// 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

// By considering the terms in the Fibonacci sequence whose values do not exceed four million,
// find the sum of the even-valued terms.

const (
	max     = 4e6
	divisor = 2
)

func main() {
	result := 0
	vals := [2]int{1, 1}

	for vals[1] <= max {
		tmp := vals[0] + vals[1]
		vals[0] = vals[1]
		vals[1] = tmp

		if vals[1]%2 == 0 {
			result += vals[1]
		}
	}
	fmt.Println(result)
}
