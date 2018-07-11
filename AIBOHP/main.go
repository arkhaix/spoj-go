package main

import "fmt"

func lcs(a string, b string) int {
	la, lb := len(a), len(b)
	dp := make([][]int, lb+1)
	for y := 0; y < lb+1; y++ {
		dp[y] = make([]int, la+1)
	}

	for y := 1; y <= lb; y++ {
		for x := 1; x <= la; x++ {
			if a[x-1] == b[y-1] {
				dp[y][x] = 1 + dp[y-1][x-1]
			} else {
				dp[y][x] = dp[y-1][x]
				if dp[y][x-1] > dp[y-1][x] {
					dp[y][x] = dp[y][x-1]
				}
			}
		}
	}

	return dp[lb][la]
}

func solve(a string) {
	var b string
	for _, r := range a {
		b = string(r) + b
	}

	fmt.Println(len(a) - lcs(a, b))
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
