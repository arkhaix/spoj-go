package main

import "fmt"

// Segment tree
type node struct {
	left       *node
	right      *node
	parent     *node
	value      int64
	leftIndex  int
	rightIndex int
}

func build(arr []int64, left int, right int, parent *node) *node {
	n := node{left: nil, right: nil, parent: parent, value: 0, leftIndex: left, rightIndex: right}

	for i := left; i <= right; i++ {
		n.value += arr[i]
	}

	if right-left > 0 {
		mid := (left + right) / 2
		n.left = build(arr, left, mid, &n)
		n.right = build(arr, mid+1, right, &n)
	}

	return &n
}

func query(left int, right int, n *node) int64 {
	if n == nil {
		return 0
	}
	if left > n.rightIndex || right < n.leftIndex {
		return 0
	}
	if n.leftIndex >= left && n.rightIndex <= right {
		return n.value
	}
	return query(left, right, n.left) + query(left, right, n.right)
}

func update(index int, value int64, n *node) {
	if n == nil {
		return
	}
	if index < n.leftIndex || index > n.rightIndex {
		return
	}
	n.value += value
	update(index, value, n.left)
	update(index, value, n.right)
}

func main() {
	var numElements int
	fmt.Scan(&numElements)
	arr := make([]int64, numElements)
	for i := 0; i < numElements; i++ {
		fmt.Scan(&arr[i])
	}

	n := build(arr, 0, numElements-1, nil)

	var numOperations int
	fmt.Scan(&numOperations)

	var operation string
	var a, b int
	for i := 0; i < numOperations; i++ {
		fmt.Scan(&operation, &a, &b)
		if operation == "q" {
			fmt.Println(query(a-1, b-1, n))
		} else {
			update(a-1, int64(b), n)
		}
	}
}
