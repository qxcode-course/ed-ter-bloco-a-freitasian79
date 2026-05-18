package main
import "fmt"

func eh_primo(x int, div int) bool {
	if x < 2 {
		return false
	}

	if x == 2 {
		return true
	}

    if div * div > x {
		return true
	}

	if x % div == 0 {
		return false
	}

	return eh_primo(x, div + 1)
}

func enesimo_primo(n int, candidato int, primosEncontrados int) int {
    if eh_primo(candidato, 2) {
        primosEncontrados++
    }

    if primosEncontrados == n {
        return candidato
    }   

    return enesimo_primo(n, candidato + 1, primosEncontrados)
}

func encontrar_primo(n int) int {
    if n <= 0 {
        return 0
    }

    return enesimo_primo(n, 2, 0)
}
func main() {
    var qtd_primos int
    
    fmt.Scan(&qtd_primos)

    lista_primos := make([]int, 0, qtd_primos)

    for i := 1; i <= qtd_primos; i++ {
        lista_primos = append(lista_primos, encontrar_primo(i))    
    }

    fmt.Print("[")

    for i, num := range lista_primos {
        fmt.Print(num)

        if i < len(lista_primos) - 1 {
            fmt.Print(", ")
        }
    }

    fmt.Print("]\n")

}
