package main
import "fmt"

func BinarySearch(slice []int, value int) int {
    comeco := 0
    fim := len(slice) - 1
    resultado := -1

    for comeco <= fim {
        meio := (comeco + fim) / 2

        if slice[meio] == value {
            resultado = meio
            comeco = meio + 1
        } else if value > slice[meio] {
            comeco = meio + 1
        } else {
            fim = meio - 1
        }

    }

    if resultado != -1 {
        return resultado
    }

    return comeco
}
func main() {
    lista_inteiros := []int{1, 2, 2, 3, 3, 3}

    fmt.Println("O numero 4 deveria estar aqui: ")
    fmt.Println(BinarySearch(lista_inteiros, 7))
}
