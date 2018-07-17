package main

import "fmt"
import "math/big"

func solve(k int) {
	var factorial big.Int
	factorial.MulRange(1, int64(k))
	fmt.Println(&factorial)
}

func main() {
	numTests := 0
	fmt.Scanln(&numTests)
	for i := 0; i < numTests; i++ {
		var k int
		fmt.Scanln(&k)
		solve(k)
	}
}
