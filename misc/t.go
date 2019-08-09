package main

import "fmt"

func main() {
	total := 100000
	rounds := 0
	current := 0

	for current < total {
		current = 1000 + (2 * current)
		rounds++
	}

	fmt.Println(rounds * 1000)
}
