package main

import (
	"fmt"
	"math"
	"time"
)

func isPrime(value int) bool {
	if value <= 1 {
		return false
	}

	end := int(math.Sqrt(float64(value)))

	for i := 2; i < end; i++ {
		if value%i == 0 {
			return false
		}
	}
	return true
}

func factorise(value int) []int {
	var factors []int

	if value == 1 {
		return append(factors, 1)
	}

	current := value
	for current > 0 {
		if isPrime(current) {
			return append(factors, current)
		}

		for i := 2; i < current; i++ {
			if current%i == 0 {
				factors = append(factors, i)
				current /= i
				break
			}
		}
	}

	return factors
}

func main() {
	start := time.Now()

	result := factorise(600851475143)

	elapsed := time.Since(start)

	fmt.Println("Result:", result[len(result)-1])
	fmt.Println("Runtime:", elapsed)
}
