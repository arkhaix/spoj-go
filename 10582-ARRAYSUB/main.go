package main

import "fmt"
import "strings"

func sliceMax(s []int) (index int, value int) {
	for i, v := range s {
		if v >= value {
			index = i
			value = v
		}
	}
	return
}

func solve(arr []int, k int) {
	curMax := -1
	curMaxIndex := -1

	lastIndex := len(arr) - k
	res := make([]int, lastIndex+1)

	for i := 0; i <= lastIndex; i++ {
		if i > curMaxIndex {
			curMaxIndex, curMax = sliceMax(arr[i : i+k])
		}
		newIndex := i + k - 1
		if arr[newIndex] >= curMax {
			curMax = arr[newIndex]
			curMaxIndex = newIndex
		}
		res[i] = curMax
	}

	s := fmt.Sprint(res)
	fmt.Println(strings.Trim(s, "[]"))
}

func main() {
	var numElements int
	fmt.Scanln(&numElements)
	arr := make([]int, numElements)
	for i := 0; i < numElements; i++ {
		var n int
		fmt.Scan(&n)
		arr[i] = n
	}
	var subSize int
	fmt.Scan(&subSize)
	solve(arr, subSize)
}
