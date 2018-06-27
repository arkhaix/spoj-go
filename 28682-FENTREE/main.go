package main

import "fmt"

// Fenwick tree
type fentree struct {
	tree []int64
}

func build(arr []int64) *fentree {
	numElements := len(arr)

	t := fentree{}
	t.tree = make([]int64, numElements)

	for i := 0; i < numElements; i++ {
		t.update(i, arr[i])
	}

	return &t
}

func (t *fentree) update(index int, value int64) {
	n := len(t.tree)
	for i := index; i < n; i |= i + 1 {
		t.tree[i] += value
	}
}

func (t *fentree) query(left int, right int) int64 {
	return t.sum(right) - t.sum(left-1)
}

func (t *fentree) sum(index int) int64 {
	var res int64
	for index > 0 {
		res += t.tree[index]
		index &= index + 1
		index = index - 1
	}
	return res
}

func main() {
	var numElements int
	fmt.Scan(&numElements)
	arr := make([]int64, numElements+1)
	for i := 1; i <= numElements; i++ {
		fmt.Scan(&arr[i])
	}

	tree := build(arr)

	var numOperations int
	fmt.Scan(&numOperations)

	var operation string
	var a, b int
	for i := 0; i < numOperations; i++ {
		fmt.Scan(&operation, &a, &b)
		if operation == "q" {
			fmt.Println(tree.query(a, b))
		} else {
			tree.update(a, int64(b))
		}
	}
}
