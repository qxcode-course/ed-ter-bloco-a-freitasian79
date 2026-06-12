package main

import "fmt"

func Sum(numbers []int, index, n, current_sum, target int) bool {
	if current_sum == target {
		return true
	}

	if current_sum > target || index == n {
		return false
	}

	return Sum(numbers, index+1, n, current_sum+numbers[index], target) ||
	Sum(numbers, index+1, n, current_sum, target)
}

func main() {
	var n, k int

	fmt.Scan(&n, &k)

	numbers := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i])
	}

	result := Sum(numbers, 0, n, 0, k)

	if result {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}