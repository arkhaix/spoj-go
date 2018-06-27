package main

import "fmt"

func solve(x int, y int) {
	if x == y {
		if x%2 == 0 {
			fmt.Println(x * 2)
		} else {
			fmt.Println(((x - 1) * 2) + 1)
		}
	} else if x == y+2 {
		if x%2 == 0 {
			fmt.Println((y * 2) + 2)
		} else {
			fmt.Println((y * 2) + 1)
		}
	} else {
		fmt.Println("No Number")
		return
	}
}

func main() {
	numTests := 0
	fmt.Scanln(&numTests)
	for i := 0; i < numTests; i++ {
		var x, y int
		fmt.Scanln(&x, &y)
		solve(x, y)
	}
}
