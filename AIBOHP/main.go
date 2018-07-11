package main

import "fmt"

type lcskey struct {
	a string
	b string
}

func makeKey(a string, b string) lcskey {
	if a < b {
		return lcskey{a, b}
	}
	return lcskey{b, a}
}

var memo map[lcskey]int

func lcs(a []rune, b []rune) int {
	key := makeKey(string(a), string(b))
	if val, ok := memo[key]; ok {
		return val
	}

	if len(a) == 0 || len(b) == 0 {
		memo[key] = 0
		return 0
	}

	if a[0] == b[0] {
		return 1 + lcs(a[1:], b[1:])
	}

	ra := lcs(a, b[1:])
	rb := lcs(a[1:], b)
	if ra > rb {
		return ra
	}
	return rb
}

func solve(k string) {
	a := []rune(k)

	var b []rune
	for i := len(a) - 1; i >= 0; i-- {
		b = append(b, a[i])
	}

	fmt.Println(len(a) - lcs(a, b))
}

func main() {
	memo = make(map[lcskey]int)

	numTests := 0
	fmt.Scanln(&numTests)
	for i := 0; i < numTests; i++ {
		var k string
		fmt.Scanln(&k)
		solve(k)
	}
}
