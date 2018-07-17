package main

import "fmt"

func solve(n int, items []int64) {
	dp := make([][]int64, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int64, n+1)
	}

	for day := 1; day <= n; day++ {
		for leftTaken := 0; leftTaken <= day; leftTaken++ {
			var left, right int64
			if leftTaken > 0 {
				left = dp[day-1][leftTaken-1] + (items[leftTaken-1] * int64(day))
			}
			if leftTaken < day {
				rightTaken := day - leftTaken
				right = dp[day-1][leftTaken] + (items[n-rightTaken] * int64(day))
			}

			dp[day][leftTaken] = left
			if right > left {
				dp[day][leftTaken] = right
			}
		}
	}

	var res int64
	for i := 0; i <= n; i++ {
		if dp[n][i] > res {
			res = dp[n][i]
		}
	}
	fmt.Println(res)
}

func main() {
	var n int
	fmt.Scan(&n)
	var items []int64
	for i := 0; i < n; i++ {
		var v int64
		fmt.Scan(&v)
		items = append(items, v)
	}
	solve(n, items)
}
