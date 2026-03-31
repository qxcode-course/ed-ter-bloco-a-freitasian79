    package main
    import "fmt"
    func main() {
        var n, e int

        fmt.Scan(&n, &e)

        fila := make([]int, n)

        pos := e - 1


        for i := 0; i < n; i++{
            fila[i] = i + 1
        }

        for len(fila) > 0    {
            fmt.Print("[ ")
            for i, v := range fila {
                if i == pos {
                    fmt.Printf("%d> ", v)
                } else {
                    fmt.Printf("%d ", v)
                }
            }
            fmt.Println("]")

            if len(fila) == 1{
                break
            }   

            proxVitima := (pos + 1) % len(fila)
            fila = append(fila[:proxVitima], fila[proxVitima+1:]...)
            pos = proxVitima % len(fila)
        }

        
    }
