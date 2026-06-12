package main
import "fmt"

type Pos struct {
    l, c int
}

func main() {
    
    found := false

    grid := [][]byte {
        {'I', 'A', 'X'},
        {'X', 'N', 'X'},
        {'X', 'X', 'X'}
    }

    var word = "IAN"

    nl := len(grid)
    nc := len(grid[0])

    dirs := []Pos{
        {}
    }

}