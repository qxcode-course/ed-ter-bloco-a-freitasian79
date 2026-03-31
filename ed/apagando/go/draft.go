package main

import "fmt"

func main() {
	var (
		n int
		m int
	)

	// 1a ENTRADA - quantidades de pessoas na fila
	fmt.Scan(&n)

	//vetor que armazena os identificadores de cada pessoa na fila
	fila := make([]int, n)

	// 2a ENTRADA - DETERMINAÇÃO DO IDENTIFICADOR ÚNICO DE CADA INTEGRANTE DA FILA
	for i := 0; i < len(fila); i++ {
		fmt.Scan(&fila[i])
	}

	// 3a QUANTAS PESSOAS SAIRAM DA FILA
	fmt.Scan(&m)

	// 4a - IDENTIFICADORES DAS PESSOAS QUE SAÍRAM DA FILA
	desistentes := make([]int, m)

	for i := 0; i < m; i++ {
		fmt.Scan(&desistentes[i])
	}

	novaFila := make([]int, 0, n-m)

	for i := range fila {
		estaNosDesistentes := false

		for j := range desistentes {
			if fila[i] == desistentes[j] {
				estaNosDesistentes = true
				break
			}

		}

		if !estaNosDesistentes {
			novaFila = append(novaFila, fila[i])
		}
	}

	for n := range novaFila {
		fmt.Printf("%d ", novaFila[n])
	}

	fmt.Println()
}
