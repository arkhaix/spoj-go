package main

import "fmt"

type node struct {
	x, y, v int
}

func visit(x int, y int, rows int, cols int, a [][]int) {
	vis := make([][]bool, rows)
	for i := range vis {
		vis[i] = make([]bool, cols)
	}
	vis[y][x] = true

	queue := []node{node{x, y, 0}}
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		if n.v < a[n.y][n.x] {
			a[n.y][n.x] = n.v
			if n.y-1 >= 0 && vis[n.y-1][n.x] == false {
				vis[n.y-1][n.x] = true
				queue = append(queue, node{n.x, n.y - 1, n.v + 1})
			}
			if n.y+1 < rows && vis[n.y+1][n.x] == false {
				vis[n.y+1][n.x] = true
				queue = append(queue, node{n.x, n.y + 1, n.v + 1})
			}
			if n.x-1 >= 0 && vis[n.y][n.x-1] == false {
				vis[n.y][n.x-1] = true
				queue = append(queue, node{n.x - 1, n.y, n.v + 1})
			}
			if n.x+1 < cols && vis[n.y][n.x+1] == false {
				vis[n.y][n.x+1] = true
				queue = append(queue, node{n.x + 1, n.y, n.v + 1})
			}
		}
	}
}

func solve(a [][]int, rows int, cols int) {
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if a[y][x] == 0 {
				a[y][x] = 1
				visit(x, y, rows, cols, a)
			}
		}
	}

	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			fmt.Print(a[y][x])
			if x < cols-1 {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}

func main() {
	numTests := 0
	fmt.Scan(&numTests)
	for i := 0; i < numTests; i++ {
		var rows, cols int
		fmt.Scan(&rows, &cols)
		a := make([][]int, rows)
		for i := range a {
			a[i] = make([]int, cols)
		}

		var s string
		for y := 0; y < rows; y++ {
			fmt.Scanln(&s)
			for x := 0; x < cols; x++ {
				if s[x] == '0' {
					a[y][x] = 1000000000
				}
			}
		}

		solve(a, rows, cols)
	}
}
