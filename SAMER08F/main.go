package main

import "fmt"

func main() {
	// Just precalculate all the answers
	sumSquares := make([]int, 101)
	sumSquares[0] = 0
	for i := 1; i <= 100; i++ {
		sumSquares[i] = sumSquares[i-1] + (i * i)
	}

	// Run the tests
	for {
		var k int
		fmt.Scanln(&k)
		if k == 0 {
			break
		} else {
			fmt.Println(sumSquares[k])
		}
	}
}
