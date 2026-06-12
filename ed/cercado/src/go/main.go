package main

import (
	"bufio"
	"fmt"
	"os"
)

func inBorder(i, j, rows, cols int) bool {
	return i == 0 || i == linhas - 1 || j == 0 || j == colunas - 1
}

// NÃO ALTERE A ASSINATURA DA FUNÇÃO solve
func solve(board [][]byte) {
	visited := make(map[int]bool)

	for i := 0 ; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if inBorder(i, j, len(board), len(board[0]) {
				visited[board[i][j]] = true
			}
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if !visited[board[i][j]] {
				board[i][j] = 'X'
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
