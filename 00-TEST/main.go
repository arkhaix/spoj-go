package main

import "fmt"

func main() {
	val := 0
	for val != 42 {
		_, _ = fmt.Scanln(&val)
		if val != 42 {
			fmt.Println(val)
		}
	}
}
