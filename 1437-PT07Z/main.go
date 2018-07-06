package main

import "fmt"

type node struct {
	next    []*node
	visited bool
}

var maxLen int
var maxNode *node
var nodes []node

func init() {
	nodes = make([]node, 10001)
}

func dfs(n *node, curLen int) {
	if n.visited == true {
		return
	}

	curLen++
	if curLen > maxLen {
		maxLen = curLen
		maxNode = n
	}

	n.visited = true
	for _, nextNode := range n.next {
		dfs(nextNode, curLen)
	}
	n.visited = false
}

func solve(start *node) {
	// Longest path from the arbitrarily-chosen starting node
	// will give us a leaf node which is guaranteed to be part
	// of the longest actual path
	dfs(start, 0)

	// maxNode now contains that farthest leaf node
	dfs(maxNode, 0)

	fmt.Println(maxLen - 1)
}

func main() {
	var start *node
	numNodes := 0
	fmt.Scan(&numNodes)
	for i := 0; i < numNodes; i++ {
		var u, v int
		fmt.Scan(&u, &v)
		nodes[u].next = append(nodes[u].next, &nodes[v])
		nodes[v].next = append(nodes[v].next, &nodes[u])
		if start == nil {
			start = &nodes[u]
		}
	}
	solve(start)
}
