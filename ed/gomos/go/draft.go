package main

import "fmt"

func main() {
	var Q int
	var d string

	fmt.Scan(&Q)

	fmt.Scan(&d)

	x := make([]int, Q)
	y := make([]int, Q)

	// contagem dos gomos
	for i := 0; i < Q; i++ {
		fmt.Scan(&x[i], &y[i])
	}

	for i := Q - 1; i > 0; i-- {
		x[i], y[i] = x[i-1], y[i-1]
	}

	//movimento da cabeça

	switch d {
	case "L":
		x[0]--
	case "R":
		x[0]++
	case "U":
		y[0]--
	case "D":
		y[0]++
	}

	for i := 0; i < Q; i++ {
		fmt.Printf("%d %d\n", x[i], y[i])
	}
}
