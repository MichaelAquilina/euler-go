package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	current := 1
	prev := 1
	total := 0

	for current < 4000000 {
		tmp := current
		current = current + prev
		prev = tmp

		if current%2 == 0 {
			total += current
		}
	}

	elapsed := time.Since(start)

	fmt.Println("Result:", total)
	fmt.Println("Runtime:", elapsed)
}
