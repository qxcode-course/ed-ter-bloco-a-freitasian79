package main
import "fmt"
func main() {
    var total, baruel_possui int

    fmt.Scan(&total)
    fmt.Scan(&baruel_possui)

    lista_baruel := make([]int, baruel_possui)
    album_baruel := make(map[int]int)

    for i := range lista_baruel {
        fmt.Scan(&lista_baruel[i])
        album_baruel[lista_baruel[i]]++
    }

    repetidas := []int{}

    for i := 1; i < baruel_possui; i++ {
        if lista_baruel[i] == lista_baruel[i - 1] {
            repetidas = append(repetidas, lista_baruel[i])
        }
    }

    faltantes := []int{}

    for i := 1; i <= total; i++ {
        if album_baruel[i] == 0 {
            faltantes = append(faltantes, i)
        }
    }

    // imprimindo os resultados

    if len(repetidas) == 0 {
        fmt.Println("N")
    } else {
        for i, v := range repetidas {
            fmt.Print(v)
            if i < len(repetidas)-1 {
                fmt.Print(" ")
            }
            
        } 
        fmt.Println()

    }

    if len(faltantes) == 0 {
        fmt.Println("N")
    } else {
        for i, v := range faltantes {
            fmt.Print(v)
            if i < len(faltantes) - 1 {
                fmt.Print(" ")
            }
        }

        fmt.Println()
    }
    
}
