package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	total := 0
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			total += i
		}
	}

	elapsed := time.Since(start)

	fmt.Println("Result:", total)
	fmt.Println("Runtime:", elapsed)
}
