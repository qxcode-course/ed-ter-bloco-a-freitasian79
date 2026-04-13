package main

import "fmt"

func exibir_lista(lista []int, tamanho int) {
    fmt.Print("[ ")
    for i := 0; i < tamanho; i++ {
        fmt.Printf("%d ", lista[i])
    }
    fmt.Println("]")
}

func main() {
    var T, R int
    fmt.Scan(&T, &R)

    lista := make([]int, T)

    for i := 0; i < T; i++ {
        fmt.Scan(&lista[i])
    }

    lista_reorganizada := make([]int, T)

    for i := 0; i < T; i++ {
        novaPos := (i + R) % T
        lista_reorganizada[novaPos] = lista[i]
    }

    exibir_lista(lista_reorganizada, T)
}