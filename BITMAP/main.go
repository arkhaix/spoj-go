package main

import "fmt"

func visit(n, x int, y int, rows int, cols int, a [][]int) {
	if x < 0 || x >= cols || y < 0 || y >= rows {
		return
	}
	if a[y][x] >= 0 && a[y][x] <= n {
		return
	}
	a[y][x] = n
	visit(n+1, x, y-1, rows, cols, a)
	visit(n+1, x, y+1, rows, cols, a)
	visit(n+1, x-1, y, rows, cols, a)
	visit(n+1, x+1, y, rows, cols, a)
}

func solve(a [][]int, rows int, cols int) {
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if a[y][x] == 0 {
				a[y][x] = 1
				visit(0, x, y, rows, cols, a)
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
					a[y][x] = -1
				}
			}
		}

		solve(a, rows, cols)
	}
}
