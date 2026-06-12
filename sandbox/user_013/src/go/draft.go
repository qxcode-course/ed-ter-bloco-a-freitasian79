package main

import "fmt"

type Pos struct {
	l, c int
}

// CORRIGIDO: Adicionado 'func' e ajustada a sintaxe do parâmetro 'visited'
func dfs(grid [][]byte, p Pos, visited map[Pos]bool) {
	// MELHORIA DE SEGURANÇA: trocado len(grid[0]) por len(grid[p.l]) para evitar panics futuros
	if p.l < 0 || p.l >= len(grid) || p.c < 0 || p.c >= len(grid[p.l]) {
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

func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	count := 0
	visited := make(map[Pos]bool)

	for l := 0; l < len(grid); l++ {
		for c := 0; c < len(grid[l]); c++ {
			if grid[l][c] == '1' && !visited[Pos{l, c}] {
				dfs(grid, Pos{l, c}, visited)
                count++
			}
		}
	}

	return count
}

func main() {
	grid := [][]byte{
		{'1', '1', '0'},
		{'1', '0', '0'},
		{'1', '0', '1'}, // CORRIGIDO: Adicionada a vírgula obrigatória aqui no final
	}

	result := numIslands(grid)
	fmt.Println(result) // Vai printar corretamente: 2
}