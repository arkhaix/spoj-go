package main

import "fmt"

func getIntervalSums(values []int, marks []int) []int {
	var res []int
	if len(marks) == 0 {
		sum := 0
		for _, v := range values {
			sum += v
		}
		res = append(res, sum)
		return res
	}

	var vi, mi int
	sum := 0
	for vi, mi = 0, 0; vi < len(values); vi++ {
		sum += values[vi]
		if mi < len(marks) && values[vi] == marks[mi] {
			res = append(res, sum)
			sum = 0
			mi++
		}
	}
	res = append(res, sum)

	return res
}

func solve(a []int, b []int) {
	aContains := make(map[int]int)
	for i, v := range a {
		aContains[v] = i
	}

	bContains := make(map[int]int)
	var shared []int
	for i, v := range b {
		bContains[v] = i
		if _, ok := aContains[v]; ok {
			shared = append(shared, v)
		}
	}

	aSums := getIntervalSums(a, shared)
	bSums := getIntervalSums(b, shared)
	// fmt.Println("shared:", shared)
	// fmt.Println("a:", aSums)
	// fmt.Println("b:", bSums)

	res := 0
	for i := 0; i < len(aSums); i++ {
		if aSums[i] > bSums[i] {
			res += aSums[i]
		} else {
			res += bSums[i]
		}
	}

	fmt.Println(res)
	// fmt.Println("")
}

func main() {
	for {
		var n int
		fmt.Scan(&n)
		if n == 0 {
			break
		}

		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&a[i])
		}

		fmt.Scan(&n)
		b := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&b[i])
		}

		solve(a, b)
	}
}
