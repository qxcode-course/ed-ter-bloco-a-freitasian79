package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func MagicSearch(slice []int, value int) int {
	// busca binaria
	start := 0
	end := len(slice) - 1
	resultado := -1

	for start <= end {
		middle := (start + end) / 2

		if slice[middle] == value {
			start = middle + 1
			resultado = middle
		} else if value > slice[middle] {
			start = middle + 1
		} else {
			end = middle - 1
		}
	}

	if resultado != -1 {
		return resultado
	} 

	return start
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	slice := make([]int, 0, 1)
	for _, elem := range parts[1 : len(parts)-1] {
		value, _ := strconv.Atoi(elem)
		slice = append(slice, value)
	}

	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	result := MagicSearch(slice, value)
	fmt.Println(result)
}
