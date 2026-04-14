package main
import "fmt"
func main() {

	var n, m int

	fmt.Scan(&n)
	na_fila := make([]int, n)

	//identificadores unicos
	for i := 0; i < len(na_fila); i++ {
		fmt.Scan(&na_fila[i])
	}

	fmt.Scan(&m)
	vai_sair := make(map[int]bool)
	for i := 0; i < m; i++ {
		var x int
		fmt.Scan(&x)
		vai_sair[x] = true
	}

	for _, pessoa := range na_fila {
		if !vai_sair[pessoa] {
			fmt.Printf("%d ", pessoa)
		}
	}
	fmt.Print("\n")

}
