package main

import (
	"bufio"
	"fmt"
	"os"
)


func exist(grid [][]byte, word string) bool {
	n_rows := len(grid)
	n_cols := len(grid[0])

	for i := 0; i < n_rows; i++ {
		for j := 0; j < n_cols; j++ {
			if backtrack(grid, n_rows, n_cols, i, j, word, 0) {
				return true
			}
		}
	}

	return false
}

func backtrack(grid [][]byte, n_rows, n_cols, i, j int, word string, index int) bool {
	if index == len(word) {
		return true
	}

	if i < 0 || i >= n_rows || j < 0 || j >= len(grid[0]) || grid[i][j] != word[index] {
		return false
	}

	temp := grid[i][j]
	grid[i][j] = '#'

	if backtrack(grid, n_rows, n_cols, i+1, j, word, index+1) ||
	backtrack(grid, n_rows, n_cols, i-1, j, word, index+1) ||
	backtrack(grid, n_rows, n_cols, i, j+1, word, index+1) ||
	backtrack(grid, n_rows, n_cols, i, j-1, word, index+1) {
		return true
	}

	grid[i][j] = temp
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var word string
	fmt.Sscanf(scanner.Text(), "%s", &word)
	grid := make([][]byte, 0)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	if exist(grid, word) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
