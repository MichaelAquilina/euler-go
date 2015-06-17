package main

import (
	"fmt"
	"strconv"
	"time"
)

func isPalindrome(value string) bool {
	strLength := len(value)
	if strLength == 0 {
		return true
	}

	if strLength%2 != 0 {
		return false
	}

	return value[0] == value[strLength-1] && isPalindrome(value[1:strLength-1])
}

func main() {
	start := time.Now()

	maxValue := 0

	for i := 100; i < 999; i++ {
		for j := 100; j < 999; j++ {
			result := i * j
			if result > maxValue && isPalindrome(strconv.Itoa(result)) {
				maxValue = result
			}
		}
	}

	runtime := time.Since(start)

	fmt.Println("Result:", maxValue)
	fmt.Println("Runtime:", runtime)
}
