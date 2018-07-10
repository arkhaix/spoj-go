package main

import "fmt"

var fives []uint64

func solve(k uint64) {
	var result uint64
	result = 0
	for i := 1; i < len(fives); i++ {
		if fives[i] > k {
			break
		}
		result += k / fives[i]
	}
	fmt.Printf("%d\n", result)
}

func prep() {
	fives = append(fives, 0)
	var val uint64
	for val = 5; val < 1000000000; val *= 5 {
		fives = append(fives, val)
	}
}

func main() {
	prep()
	numTests := 0
	fmt.Scanln(&numTests)
	for i := 0; i < numTests; i++ {
		var k uint64
		fmt.Scanln(&k)
		solve(k)
	}
}
