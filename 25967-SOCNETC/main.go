package main

import "fmt"

var parent []int
var size []int
var maxSize int

func root(x int) int {
	for parent[x] != x {
		x = parent[x]
	}
	return x
}

func merge(x, y int) {
	rx := root(x)
	ry := root(y)
	if size[rx]+size[ry] <= maxSize {
		if size[rx] < size[ry] {
			parent[rx] = ry
			size[ry] = size[ry] + size[rx]
		} else {
			parent[ry] = rx
			size[rx] = size[rx] + size[ry]
		}
	}
}

func query(x, y int) {
	if root(x) == root(y) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func getSize(x int) {
	fmt.Println(size[root(x)])
}

func main() {
	var numUsers int
	fmt.Scan(&numUsers, &maxSize)
	numUsers = numUsers + 1

	parent = make([]int, numUsers)
	size = make([]int, numUsers)
	for i := 0; i < numUsers; i++ {
		parent[i] = i
		size[i] = 1
	}

	var numQueries int
	fmt.Scan(&numQueries)

	for i := 0; i < numQueries; i++ {
		var command string
		fmt.Scan(&command)

		if command == "A" {
			var x, y int
			fmt.Scan(&x, &y)
			merge(x, y)
		} else if command == "E" {
			var x, y int
			fmt.Scan(&x, &y)
			query(x, y)
		} else /* S */ {
			var x int
			fmt.Scan(&x)
			getSize(x)
		}
	}
}
