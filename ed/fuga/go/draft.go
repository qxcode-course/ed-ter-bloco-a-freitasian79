package main

import "fmt"

func main() {
	var (
		h int
		p int
		f int
		d int
	)

	fmt.Scan(&h, &p, &f, &d)

	var aux string

	//sentido horário (esquerda)
	if d == -1 {
        for {
            if f == p {
                aux = "N"
                break
            }

            if f == h {
                aux = "S"
                break
            }
            
            f--
            if f < 0 {
                f = 15
            }
         }
    } else {
        for {
            if f == p {
                aux = "N"
                break
            }

            if f == h {
                aux = "S"
                break
            }

            f++

            if f > 15 {
                f = 0
            }
        }
    }

    fmt.Println(aux)

}
