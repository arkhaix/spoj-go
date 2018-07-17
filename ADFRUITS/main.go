package main

import "fmt"

func solve(a, b string) {
	la, lb := len(a), len(b)
	dp := make([][]int, lb+1)
	for y := 0; y < lb+1; y++ {
		dp[y] = make([]int, la+1)
	}

	// Find the LCS
	for y := 1; y < lb+1; y++ {
		for x := 1; x < la+1; x++ {
			if a[x-1] == b[y-1] {
				dp[y][x] = dp[y-1][x-1] + 1
			} else {
				dp[y][x] = dp[y-1][x]
				if dp[y-1][x] < dp[y][x-1] {
					dp[y][x] = dp[y][x-1]
				}
			}
		}
	}

	// Get the actual LCS string by walking
	// the lcs dp table in reverse
	lcs := ""
	for x, y := la, lb; x != 0 && y != 0; {
		if a[x-1] == b[y-1] {
			lcs = string(a[x-1]) + lcs
			y--
			x--
		} else {
			if dp[y-1][x] > dp[y][x-1] {
				y--
			} else {
				x--
			}
		}
	}

	// Print the two strings out character-by-character,
	// joining them up at the shared characters
	// (indicated by being in lcs)
	res := ""
	var ia, ib int
	for _, v := range lcs {
		for ia < la && rune(a[ia]) != v {
			res += string(a[ia])
			ia++
		}
		for ib < lb && rune(b[ib]) != v {
			res += string(b[ib])
			ib++
		}
		res += string(v)
		ia++
		ib++
	}
	for ia < la {
		res = res + string(a[ia])
		ia++
	}
	for ib < lb {
		res = res + string(b[ib])
		ib++
	}

	fmt.Println(res)
}

func main() {
	var err error
	for err == nil {
		var a, b string
		_, err = fmt.Scan(&a, &b)
		if err == nil {
			solve(a, b)
		}
	}
}
