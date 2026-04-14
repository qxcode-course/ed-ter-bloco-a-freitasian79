package main
import "fmt"
func main() {
    qtd := 0
    fmt.Scan(&qtd)

    animais := make([]int, qtd)
    solteiros := make(map[int]int)

    for i := range animais {
        fmt.Scan(&animais[i])
    }

    casais := 0
    for _, animal := range animais {    
        par_procurado := -animal

        if solteiros[par_procurado] > 0 {
            casais++
            solteiros[par_procurado]--
        } else {
            solteiros[animal]++
        }
    }

    fmt.Println(casais)
}
