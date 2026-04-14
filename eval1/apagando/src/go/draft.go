package main
import "fmt"
func main() {
	var n int

	fmt.Scan(&n)
	na_fila := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&na_fila[i])
	}

	var m int
	fmt.Scan(&m)
	vai_sair := make(map[int]bool)

	for i := 0; i < m; i++ {
		var sai int 
		fmt.Scan(&sai)
		vai_sair[sai] = true
	}


	for _, pessoa := range na_fila {
		if !vai_sair[pessoa] {
			fmt.Printf("%d ", pessoa)
		}
	}

	fmt.Println()


}
