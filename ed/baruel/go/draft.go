package main

import "fmt"

func main() {
	var totalAlbum int
	var figsBaruel int
	fmt.Scan(&totalAlbum)
	fmt.Scan(&figsBaruel)

	figurinhas := make([]int, figsBaruel)

	for i := 0; i < figsBaruel; i++ {
		fmt.Scan(&figurinhas[i])
	}

	repetidas := []int{}
    faltantes := []int{}

    for i := 0; i < figsBaruel; i++ {
        if i > 0 && figurinhas[i] == figurinhas[i - 1] {
            repetidas = append(repetidas, figurinhas[i])
        }
    }

    // procurar as faltantes

    for num := 1; num <= totalAlbum; num++ {
        temNaMao := false

        for j := 0; j < figsBaruel; j++ {
            if figurinhas[j] == num {
                temNaMao = true
                break
            }
        }

        if !temNaMao {
            faltantes = append(faltantes, num)
        }
    }

    if len(repetidas) == 0 {
        fmt.Println("N")
    } else {
        for i := 0; i < len(repetidas); i++ {
            fmt.Print(repetidas[i])
            if i < len(repetidas)-1 {
                fmt.Print(" ")
            }
        }
        fmt.Println()
    }

    
    if len(faltantes) == 0{
        fmt.Println("N")
    } else {
        for i := 0; i < len(faltantes); i++ {
            fmt.Print(faltantes[i])
            if i < len(faltantes) - 1 {
                fmt.Print(" ")
            }
        }
        fmt.Println()
    }
    
    
}
