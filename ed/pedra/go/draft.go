package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		numeroCompetidores int
		a                  int
		b                  int
	)

	fmt.Scan(&numeroCompetidores)

	listaCompetidores := make([]int, numeroCompetidores)

	for i := 0; i < numeroCompetidores; i++ {
		fmt.Scan(&a, &b)
		if a < 10 || b < 10 {
			listaCompetidores[i] = -1 //valor que será atribuido para os jogadores desclassificados
		} else {
			listaCompetidores[i] = int(math.Abs(float64(a - b)))
		}

	}
	var menor int = 999
	var indiceVencedor int = 0
	for e := 0; e < len(listaCompetidores); e++ {
		if listaCompetidores[e] != -1 && listaCompetidores[e] < menor {
			menor = listaCompetidores[e]
			indiceVencedor = e
		}
	}

	if menor == 999 {
		fmt.Println("sem ganhador")
	} else {
		fmt.Println(indiceVencedor)
	}

}
