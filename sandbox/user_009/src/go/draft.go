package main
import "fmt"

type Pos struct {
    l, c int
}

func getNeig(p Pos) []Pos {
    return []Pos{{p.l, p.c + 1}, {p.l, p.c -1}, {p.l - 1, p.c}, {p.l + 1, p.c}}
}

func inside(grid [][]rune, p Pos) bool {
    return !(p.l < 0 || p.l >= len(grid) || p.c < 0 || p.c >= len(grid[0])) 
}

func match(grid [][]rune, p Pos, value rune) bool {
    return inside(grid, p) && grid[p.l][p.c] == value
}

func search(grid [][]rune, endPos, p Pos, visited map[Pos]bool) bool {
    if !match(grid, p, '.') || visited[p] {
        return false
    }

    visited[p] = true

    if p == endPos {
        grid[p.l][p.c] = '*'
        return true
    }

    for _, elem := range getNeig(p) {
        if search(grid, endPos, elem, visited) {
            grid[p.l][p.c] = '*'
            return true
        }
    }

    return false
}

func printGrid(grid [][]rune) {

    for _, line := range grid {
        fmt.Println(string(line))
    }
}
func main() {
    
    grid := [][]rune{
		[]rune("...#."),
		[]rune(".#..."),
		[]rune("..#.."),
		[]rune("....."),
	}
    
    printGrid(grid)
    fmt.Println()
    startPos, endPos := Pos{0,0}, Pos{2,3}
    visited := make(map[Pos]bool, 0)
    search(grid, endPos, startPos, visited)
    printGrid(grid)
}