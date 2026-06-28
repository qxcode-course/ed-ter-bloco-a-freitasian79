package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	l, c int
}

func inBorder(i, j, rows, cols int,) bool {
	return i == 0 || i == rows - 1 || j == 0 || j == cols - 1
}

func dfs(matriz [][]byte, i, j int, visited map[Point]bool) {
	rows := len(matriz)
	cols := len(matriz[0])

	p := Point{i, j}

	if i < 0 || i >= rows || j < 0 || j >= cols || matriz[i][j] != 'O' || visited[p] {
		return
	}

	visited[p] = true 

	dfs(matriz, i+1, j, visited)
	dfs(matriz, i-1, j, visited)
	dfs(matriz, i, j+1, visited)
	dfs(matriz, i, j-1, visited)
}

// NÃO ALTERE A ASSINATURA DA FUNÇÃO solve
func solve(board [][]byte) {
	visited := make(map[Point]bool)

	rows := len(board)
	cols := len(board[0])

	for i := 0 ; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if inBorder(i, j, rows, cols) && board[i][j] == 'O' {
				dfs(board, i, j, visited)
			}
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'O'{
				pos := Point{i, j}

				if !visited[pos] {
					board[i][j] = 'X'
				}
			}
		}
	}

}

// NÃO ALTERE A MAIN
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var nrows, ncols int
	fmt.Sscanf(scanner.Text(), "%d %d", &nrows, &ncols)
	board := make([][]byte, nrows)
	for i := 0; i < nrows; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}
	solve(board)
	for _, row := range board {
		fmt.Println(string(row))
	}
}
