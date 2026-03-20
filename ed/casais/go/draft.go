package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	animais := make([]int, n)

	var totalCasais int
	for i := 0; i < n; i++ {
		fmt.Scan(&animais[i])
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if animais[i] > 0 || animais[i] < 0 {
				if animais[i] == animais[j]*-1 {
					animais[j], animais[i] = 0, 0
					totalCasais++
				}
			}
		}
	}

	fmt.Println(totalCasais)

}


