package main

import "fmt"

func dfs(grid [][]byte, i, j int) {
	rows := len(grid)
	cols := len(grid[0])

	if i < 0 || j < 0 || i >= rows || j >= cols {
		return
	}

	fmt.Println(string(grid[i][j]))

	dfs(grid, i+1, j)
	dfs(grid, i-1, j)
	dfs(grid, i, j+1)
	dfs(grid, i, j-1)
}

func main() {
	grid := [][]byte{
		{'A', 'B'},
		{'C', 'D'},
	}

	dfs(grid, 0, 0)
}