package main

import "fmt"

/* func main() {
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
    
    
} */

func main(){
    var total_album, baruel_possui int

    fmt.Scan(&total_album)
    fmt.Scan(&baruel_possui)

    possui := make([]int, baruel_possui)
    album := make(map[int]int)

    for i := 0; i < baruel_possui; i++ {
        fmt.Scan(&possui[i])
        album[possui[i]]++
    }

    repetidas := []int{}
    for i := 1; i < baruel_possui; i++ {
        if possui[i] == possui[i - 1] {
            repetidas = append(repetidas, possui[i])
        }
    }

    faltantes := []int{}
    for i := 1; i <= total_album; i++ {
        if album[i] == 0 {
            faltantes = append(faltantes, i)
        }
    }

    // imprimindo resultados

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
            if i < len(faltantes)-1 {
                fmt.Print(" ")
            }
        }
        fmt.Println()
    }
} 
