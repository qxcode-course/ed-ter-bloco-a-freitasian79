package main
import "fmt"

func busca_binaria(lista []int, alvo int) int{
    start := 0
    end := len(lista) - 1

    for start <= end {
        mid := (start + end) / 2

        if lista[mid] == alvo {
            return mid
        }

        if lista[mid] < alvo {
            start = mid + 1 
        } else {
            end = mid - 1
        }
    }

    return -1
}
func main() {
    lista := []int{10, 20, 30, 40, 50   }

    var alvo int

    fmt.Println("Digite o valor alvo que deseja encontrar: ")
    fmt.Scan(&alvo)

    fmt.Println(lista)

    resultado := busca_binaria(lista, alvo)

    if resultado != -1 {
        fmt.Println(resultado)
    }

}
