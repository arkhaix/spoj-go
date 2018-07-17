package main

import "fmt"

func solve(k string) {

}

func main() {
	numTests := 0
	fmt.Scanln(&numTests)
	for i := 0; i < numTests; i++ {
		var k string
		fmt.Scanln(&k)
		solve(k)
	}
}
