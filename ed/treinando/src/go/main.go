package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func tostr(vet []int) string {
	if len(vet) == 0 {
		return "[]"
	}

	if len(vet) == 1 {
		return "[" + strconv.Itoa(vet[0]) + "]"
	}

	resto := tostr(vet[1:])

	return "[" + strconv.Itoa(vet[0]) + ", " + resto[1:]
}

func tostrrev(vet []int) string {
	if len(vet) == 0 {
		return "[]"
	}

	if len(vet) == 1 {
		return "[" + strconv.Itoa(vet[0]) + "]"
	}

	restante := tostrrev(vet[1:])

	return restante[:len(restante) - 1] + ", " + strconv.Itoa(vet[0]) + "]"
}

// reverse: inverte os elementos do slice
func reverse(vet []int) {
	if len(vet) <= 1 {
		return 
	}
	ultimo := len(vet) - 1
	vet[0], vet[ultimo] = vet[ultimo], vet[0]  
	reverse(vet[1 : ultimo])
}

// sum: soma dos elementos do slice
func sum(vet []int) int {
	if len(vet) == 0{
		return 0
	}

	return vet[0] + sum(vet[1:])
}

// mult: produto dos elementos do slice
func mult(vet []int) int {
	if len(vet) == 0{
		return 1
	}

	return vet[0] * mult(vet[1:])
}

// min: retorna o índice e valor do menor valor
// crie uma função recursiva interna do modelo
// var rec func(v []int) (int, int)
// para fazer uma recursão que retorna valor e índice
func min(vet []int) int {
	if len(vet) < 1 {
		return -1
	}

	if len(vet) == 1 {
		return 0
	}

	menor_resto := min(vet[1:])

	indice_resto := menor_resto + 1

	if vet[0] <= vet[indice_resto] {
		return 0
	}

	return indice_resto
}

func main() {
	var vet []int
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Fields(line)
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			return
		case "read":
			vet = nil
			for _, arg := range args[1:] {
				if val, err := strconv.Atoi(arg); err == nil {
					vet = append(vet, val)
				}
			}
		case "tostr":
			fmt.Println(tostr(vet))
		case "torev":
			fmt.Println(tostrrev(vet))
		case "reverse":
			reverse(vet)
		case "sum":
			fmt.Println(sum(vet))
		case "mult":
			fmt.Println(mult(vet))
		case "min":
			fmt.Println(min(vet))
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
