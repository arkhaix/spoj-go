package main

import "fmt"

var dp []int
var numDp = 10000000

func init() {
	dp = make([]int, numDp)
	dp[0] = 0
	dp[1] = 1

	for i := 2; i < numDp; i++ {
		n := dp[i/2] + dp[i/3] + dp[i/4]
		if i > n {
			n = i
		}
		dp[i] = n
	}
}

func solve(n int) int {
	if n < numDp {
		return dp[n]
	}

	res := solve(n/2) + solve(n/3) + solve(n/4)
	if n > res {
		res = n
	}

	return res
}

func main() {
	var n int
	for {
		_, err := fmt.Scan(&n)
		if err != nil {
			break
		}
		fmt.Println(solve(n))
	}
}
