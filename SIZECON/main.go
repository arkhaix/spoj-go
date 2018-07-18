package main

import . "fmt"

func main() {
	var i, n, r int
	Scan(&i)
	for i > 0 {
		Scan(&n)
		if n > 0 {
			r += n
		}
		i--
	}
	Print(r)
}
