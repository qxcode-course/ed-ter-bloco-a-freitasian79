package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	} 

	count := 0
	visited := make(map[Pos]bool)

	for l := 0; l < len(grid); l++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[l][c] == '1' && !visited[Pos{l, c}] {
				dfs(grid, Pos{l, c}, visited)
				count++
			}
			
		}
	}

	return count
}

func dfs (grid [][]byte, p Pos, visited map[Pos]bool) {
	if p.l < 0 || p.l >= len(grid) || p.c < 0 || p.c >= len(grid[0]) {
		return
	}

	if grid[p.l][p.c] == '0' {
		return
	}

	if visited[p] {
		return
	}

	visited[p] = true

	dfs(grid, Pos{p.l - 1, p.c}, visited) 
	dfs(grid, Pos{p.l + 1, p.c}, visited) 
	dfs(grid, Pos{p.l, p.c - 1}, visited) 
	dfs(grid, Pos{p.l, p.c + 1}, visited)
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)
	grid := make([][]byte, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		grid[i] = []byte(scanner.Text())
	}
	result := numIslands(grid)
	fmt.Println(result)
}
