package main

import "fmt"

var primeFactors []int

func generatePrimes() {
	var isPrime [32000]bool
	for i := 0; i < 32000; i++ {
		isPrime[i] = true
	}

	for i := 2; i < 32000; i++ {
		if isPrime[i] == true {
			for j := i + i; j < 32000; j += i {
				isPrime[j] = false
			}
		}
	}

	for i := 2; i < 32000; i++ {
		if isPrime[i] == true {
			primeFactors = append(primeFactors, i)
		}
	}
}

func main() {
	generatePrimes()
	numTests := 0
	fmt.Scanln(&numTests)
	for i := 0; i < numTests; i++ {
		var low, hi int
		fmt.Scanln(&low, &hi)
		sieve(low, hi)
		fmt.Println("")
	}
}

func sieve(low int, hi int) {
	isPrime := make([]bool, hi-low+1)
	for i, _ := range isPrime {
		isPrime[i] = true
	}

	// 1 is not prime, special case
	if low <= 1 {
		isPrime[1-low] = false
	}

	for _, factor := range primeFactors {
		notPrime := int(low/factor) * factor
		for notPrime < low || notPrime == factor {
			notPrime += factor
		}
		for notPrime <= hi {
			isPrime[notPrime-low] = false
			notPrime += factor
		}
	}

	for i, v := range isPrime {
		if v == true {
			fmt.Println(low + i)
		}
	}
}
