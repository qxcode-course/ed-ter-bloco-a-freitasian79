package main
import "fmt"
func exibir_jogadores(jogadores []int, espada int) {
    fmt.Printf("[ ")
    for i, elem := range jogadores {
        if i == espada {
            fmt.Printf("%d> ", elem)
        } else {
            fmt.Printf("%d ", elem)
        }
    }
    fmt.Printf("]\n")
}
func main() {
    var n, e int

    fmt.Scan(&n, &e)

    jogadores := make([]int, 0, n)

    for i := 1; i <= n; i++ {
        jogadores = append(jogadores, i)
    }

    espada := e - 1

    for len(jogadores) > 1 {
        exibir_jogadores(jogadores, espada)
        
        vai_morrer := (espada + 1) % len(jogadores)
        jogadores = append(jogadores[:vai_morrer], jogadores[vai_morrer + 1:]...)
        espada = vai_morrer % len(jogadores)
    }
    
    exibir_jogadores(jogadores, espada)
}
