package main

import "fmt"

type node struct {
	v      int
	l, r   *node
	li, ri int
}

func build(a []int, i int, j int) *node {
	n := node{}
	n.li = i
	n.ri = j
	for m := i; m <= j; m++ {
		if a[m] > n.v {
			n.v = a[m]
		}
	}
	if j-i > 0 {
		mid := (i + j) / 2
		n.l = build(a, i, mid)
		n.r = build(a, mid+1, j)
	}
	return &n
}

func solve(x, y int, n *node) int {
	if y < n.li || x > n.ri {
		return -1
	}
	if n.li >= x && n.ri <= y {
		return n.v
	}
	a := solve(x, y, n.l)
	b := solve(x, y, n.r)
	if a > b {
		return a
	}
	return b
}

func main() {
	var n int
	fmt.Scanln(&n)
	a := make([]int, n+1)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i+1])
	}
	tree := build(a, 0, n)
	var numQueries int
	fmt.Scan(&numQueries)
	for i := 0; i < numQueries; i++ {
		var x, y int
		fmt.Scanln(&x, &y)
		res := solve(x, y, tree)
		fmt.Println(res)
	}
}
