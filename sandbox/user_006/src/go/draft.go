package main
import "fmt"

type Pos struct {
    l, c int
}

func exist(grid [][]byte, word string) bool {

	for l := 0; l < len(grid); l++ {
		for c := 0; c < len(grid[0]); c++ {

            visited := make(map[Pos]bool)

            if dfs(grid, Pos{l, c}, word, 0, visited) {
                return true
            }
		}
	}

	return false
}

func dfs(grid [][]byte, p Pos, word string, idx int, visited map[Pos]bool) bool {

    if idx == len(word) {
        return true
    }

    if p.l < 0 || p.l >= len(grid) || p.c < 0 || p.c >= len(grid[0]) {
        return false
    }

    if grid[p.l][p.c] != word[idx] {
        return false
    }

    if visited[p] {
        return false
    }

    visited[p] = true

    found := dfs(grid, Pos{p.l - 1, p.c}, word, idx+1, visited) ||
		dfs(grid, Pos{p.l + 1, p.c}, word, idx+1, visited) ||
		dfs(grid, Pos{p.l, p.c - 1}, word, idx+1, visited) ||
		dfs(grid, Pos{p.l, p.c + 1}, word, idx+1, visited)

    visited[p] = false

    return found
}

func main() {

	grid := [][]byte{
		{'A', 'B', 'C'},
		{'D', 'E', 'F'},
		{'G', 'H', 'I'},
	}


	fmt.Println(exist(grid, "AB"))
}
