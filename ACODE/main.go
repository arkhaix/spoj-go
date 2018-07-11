package main

import "flag"
import "fmt"

var debug *bool

func init() {
	debug = flag.Bool("debug", false, "debug")
	flag.Parse()
}

func solve(n string) {

	var digits []int

	for _, c := range n {
		digits = append(digits, int(c-'0'))
	}

	numDigits := len(digits)

	numOnes := make([]int64, numDigits)
	numTwos := make([]int64, numDigits)
	numOthers := make([]int64, numDigits)
	sums := make([]int64, numDigits)
	if digits[0] == 1 {
		numOnes[0] = 1
	} else if digits[0] == 2 {
		numTwos[0] = 1
	} else {
		numOthers[0] = 1
	}
	sums[0] = 1

	for i := 1; i < numDigits; i++ {
		digit := digits[i]
		if i < numDigits-1 {
			if digits[i+1] == 0 {
				digit *= 10
			}
		}
		if digit == 0 || digit > 9 {
			numOthers[i] = numOnes[i-1] + numTwos[i-1] + numOthers[i-1]
		} else if digit == 1 {
			numOnes[i] = numOthers[i-1] + numOnes[i-1] + numTwos[i-1]
			numOthers[i] = numOnes[i-1] + numTwos[i-1]
		} else if digit == 2 {
			numTwos[i] = numOthers[i-1] + numOnes[i-1] + numTwos[i-1]
			numOthers[i] = numOnes[i-1] + numTwos[i-1]
		} else if digit < 7 {
			numOthers[i] = numOthers[i-1] + (2 * numOnes[i-1]) + (2 * numTwos[i-1])
		} else {
			numOthers[i] = numOthers[i-1] + (2 * numOnes[i-1]) + numTwos[i-1]
		}

		sums[i] = numOnes[i] + numTwos[i] + numOthers[i]
	}
	if *debug {
		fmt.Println("Digs:", digits)
		fmt.Println("Ones:", numOnes)
		fmt.Println("Twos:", numTwos)
		fmt.Println("Oths:", numOthers)
		fmt.Println("Sums:", sums)
	}

	resIndex := numDigits - 1
	res := numOnes[resIndex] + numTwos[resIndex] + numOthers[resIndex]
	fmt.Println(res)
}

func main() {
	var n string
	fmt.Scan(&n)
	for n != "0" {
		solve(n)
		fmt.Scan(&n)
	}
}
