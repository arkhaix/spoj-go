package main

import "fmt"

func solve(h int, w int, values [][]int) {
	hh := h + 1
	ww := w + 2
	dp := make([][]int, hh)

	dp[0] = make([]int, ww)
	for y := 1; y < hh; y++ {
		dp[y] = make([]int, ww)
		for x := 1; x < ww-1; x++ {
			//mp := max(dp[y-1][x-1], dp[y-1][x], dp[y-1][x+1])
			mp := dp[y-1][x-1]
			if dp[y-1][x] > mp {
				mp = dp[y-1][x]
			}
			if dp[y-1][x+1] > mp {
				mp = dp[y-1][x+1]
			}

			dp[y][x] = values[y-1][x-1] + mp
		}
	}

	res := 0
	for _, v := range dp[hh-1] {
		if v > res {
			res = v
		}
	}

	fmt.Println(res)
}

func main() {
	numTests := 0
	fmt.Scan(&numTests)
	for i := 0; i < numTests; i++ {
		var h, w int
		fmt.Scan(&h, &w)

		values := make([][]int, h)
		for y := 0; y < h; y++ {
			values[y] = make([]int, w)
			for x := 0; x < w; x++ {
				fmt.Scan(&values[y][x])
			}
		}

		solve(h, w, values)
	}
}
