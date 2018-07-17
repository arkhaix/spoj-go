package main

import "fmt"
import "strings"

func reverse(s string) string {
	length := len(s)
	runes := make([]rune, length)
	for i, r := range s {
		runes[length-1-i] = r
	}

	return string(runes)
}

func solve(a string, b string) {
	// Get the actual numbers
	a = reverse(a)
	b = reverse(b)
	var ia, ib int
	fmt.Sscan(a, &ia)
	fmt.Sscan(b, &ib)

	// Add the numbers and print the reversed result
	result := fmt.Sprintf("%d", ia+ib)
	result = reverse(result)
	result = strings.TrimLeft(result, "0")
	fmt.Println(result)
}

func main() {
	numTests := 0
	fmt.Scanln(&numTests)
	for i := 0; i < numTests; i++ {
		var a, b string
		fmt.Scanln(&a, &b)
		solve(a, b)
	}
}
