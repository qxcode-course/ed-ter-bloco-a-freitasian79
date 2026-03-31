package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getMen(vet []int) []int {
	listaHomens := []int{}

	for i := 0; i < len(vet); i++ {
		if vet[i] > 0 {
			listaHomens = append(listaHomens, vet[i])
		}
	}

	return(listaHomens)

}

func getCalmWomen(vet []int) []int {
	listaMulheresCalmas := []int{}

	for i := 0; i < len(vet); i++{
		if vet[i] < 0 && vet[i] > -10 {
			listaMulheresCalmas = append(listaMulheresCalmas, vet[i])
		}
	}

	return listaMulheresCalmas
}

func sortVet(vet []int) []int {
	listaOrganizada := make([]int, len(vet))
	copy(listaOrganizada, vet)

	for i := 0; i < len(vet); i++ {
		for j := 0; j < len(vet) - i - 1; j++ {
			if listaOrganizada[j] > listaOrganizada[j + 1] {
				listaOrganizada[j], listaOrganizada[j + 1] = listaOrganizada[j + 1], listaOrganizada[j]
			}
		}
	}

	return listaOrganizada
}

func sortStress(vet []int) []int {
	listaOrganizadaStress := make([]int, len(vet))
	copy(listaOrganizadaStress, vet)

	tam := len(listaOrganizadaStress)

	abs := func(n int) int {
		if n < 0 {
			return -n
		}

		return n
	}

	for i := 0; i < len(listaOrganizadaStress); i++ {
		for j := 0; j < tam - i - 1; j++ {
			if abs(listaOrganizadaStress[j]) > abs(listaOrganizadaStress[j + 1]){
				listaOrganizadaStress[j], listaOrganizadaStress[j + 1] = listaOrganizadaStress[j + 1], listaOrganizadaStress[j] 
			}
		}
	}

	return listaOrganizadaStress
}

func reverse(vet []int) []int {
	listaReversa := []int{}

	for i := len(vet) - 1; i >= 0; i-- {
		listaReversa = append(listaReversa, vet[i])
	}

	return listaReversa
}
func unique(vet []int) []int {
	listaUnica := []int{}

	for i := 0; i < len(vet); i++ {
		apareceu := false

		for j := 0; j < i; j++ {
			if vet[i] == vet[j] {
				apareceu = true
				break
			}
		}

		if !apareceu {
			listaUnica = append(listaUnica, vet[i])
		}
	}

	return listaUnica
}

func repeated(vet []int) []int {
	listaRepetidos := []int{}

	for i := 0; i < len(vet); i++ {
		achou := false

		for j := 0; j < i; j++ {
			if vet[i] == vet[j]{
				achou = true
				break
			}
		}

		if achou {
			listaRepetidos = append(listaRepetidos, vet[i])
		}
	}

	return listaRepetidos
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}
