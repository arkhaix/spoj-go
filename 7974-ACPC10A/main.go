package main

import "fmt"

func solve(a int, b int, c int) {
	apc := b + (b - a)
	if c == apc {
		fmt.Printf("AP %d\n", c+(b-a))
	} else {
		fmt.Printf("GP %d\n", c*(b/a))
	}
}

func main() {
	var a, b, c int
	for {
		fmt.Scanln(&a, &b, &c)
		if a == 0 && b == 0 {
			break
		}
		solve(a, b, c)
	}
}
