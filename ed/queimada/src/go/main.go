	package main

	import (
		"bufio"
		"fmt"
		"os"
	)

	type Pos struct {
		l, c int
	}

	func getNeig(p Pos) []Pos {
		return []Pos{{p.l, p.c - 1}, {p.l - 1, p.c}, {p.l, p.c + 1}, {p.l + 1, p.c}}
	}

	func inside(grid [][]rune, p Pos) bool {
		return !(p.l < 0 || p.l >= len(grid) || p.c < 0 || p.c >= len(grid[0]))
	}

	func match(grid [][]rune, p Pos, value rune) bool {	
		return inside(grid, p) && grid[p.l][p.c] == value
	}

	func burnTrees(grid [][]rune, l, c int, visited map[Pos]bool) {
		pos := Pos{l:l, c:c}

		if !match(grid, pos, '#') || visited[pos]{
			return
		}

		visited[pos] = true
		for _, elem := range getNeig(pos) {
			burnTrees(grid, elem.l, elem.c, visited)
		}
	}

	func main() {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		line := scanner.Text()
		var nl, nc, lfire, cfire int
		fmt.Sscanf(line, "%d %d %d %d", &nl, &nc, &lfire, &cfire)

		grid := make([][]rune, 0, nl)
		for range nl {
			scanner.Scan()
			line := []rune(scanner.Text())
			grid = append(grid, line)
		}

		visited := make(map[Pos]bool, 0)
		burnTrees(grid, lfire, cfire, visited)
		for pos := range visited {
			grid[pos.l][pos.c] = 'o'
		}
		showGrid(grid)
	}

	func showGrid(grid [][]rune) {
		for _, line := range grid {
			fmt.Println(string(line))
		}
	}
