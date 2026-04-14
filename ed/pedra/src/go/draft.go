package main
import "fmt"
func absoluto(n int) int{
    if n < 0 {
        return -n
    }

    return n
}
func main() {
    var num_competidores, a, b int

    fmt.Scan(&num_competidores)
    competidores := make([]int, num_competidores)

    for i := 0; i < num_competidores; i++ {
        fmt.Scan(&a, &b)
        if a < 10 || b < 10 {
            competidores[i] = -1
        } else {
            competidores[i] = absoluto(a - b)
        }
    }

    menor := 999
    indice_vencedor := -1
    for i := 0; i < num_competidores; i++ {
        if competidores[i] != -1 {
            if competidores[i] < menor {
                menor = competidores[i]
                indice_vencedor = i
            }
        }
    }

    if indice_vencedor == -1 {
        fmt.Println("sem ganhador")
    } else {
        fmt.Println(indice_vencedor)
    }
    
}
