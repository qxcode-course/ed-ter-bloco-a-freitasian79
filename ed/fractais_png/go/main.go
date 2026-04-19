package main

import (
	"fmt"
	"math/rand"
)

func randInt(inf, sup int) float64 {
	return float64(rand.Intn(sup-inf+1) + inf)
}

/* func arvore(pen *Pen, dist float64) {
	if dist < 20 {
		return
	}	
	
	pen.SetLineWidth(dist / 8)
	pen.Walk(dist)
	pen.Right(20)
	circulo(pen, dist * 0.8)
	pen.Right(-2 * 20)
	circulo(pen, dist * 0.8)
	pen.Right(20)
	pen.Walk(-dist)
} */

// circulos => fiz tudo!
/* func circulos(pen *Pen, raio float64) {
	if raio < 5 {
		return
	}
	
	pen.DrawCircle(raio)

	for range 6 {
		pen.Right(60)
		pen.Up()
		pen.Walk(raio)
		pen.Down()

		circulos(pen, raio / 3)

		pen.Up()
		pen.Walk(-raio)
		pen.Down()

	}


} */

// floco de neve => fiz tudo
/* func floco_neve(pen *Pen, dist float64) {
	//vira direita, anda, sobe a caneta, vira mais direita, anda...
	if dist < 10 {
		return
	}

	for range 5 {
		pen.Right(72)
		pen.Walk(dist)
		floco_neve(pen, dist / 2.6)

		pen.Up()
		pen.Walk(-dist)
		pen.Down()
	}
	
} */

/* func quadrados(pen *Pen, dist float64) {
    if dist < 50 { 
        return
    }

    for range 4 {
        pen.Walk(dist)

        
        subDist := dist / 2
        recuo := subDist / 2

        pen.Up()
        
        pen.Walk(-recuo) 
        pen.Left(90)
        pen.Walk(recuo)
        pen.Right(90)
        pen.Down()

        
        quadrados(pen, subDist)

        
        pen.Up()
        pen.Right(90)
        pen.Walk(recuo)
        pen.Left(90)
        pen.Walk(recuo)
        pen.Down()
    

        pen.Right(90) 
    }
} */

// rotações => fiz tudo
/* func rotacoes(pen *Pen, dist float64) {
	if dist < 60{
		return
	}	

	pen.Walk(dist)
	pen.SetRGB(randInt(0, 255), randInt(0, 255), randInt(0, 255))
	dist -= 15
	pen.Right(90)
	rotacoes(pen, dist)
} */

/* func triangulos(pen *Pen, dist float64) {
	for range 3 {
		pen.Walk(dist)
		pen.Right(120)
	}
} */

// trigo => fiz tudo
/* func trigo(pen *Pen, dist float64){

	if dist < 10 {
		return
	}
	passo := dist / 4
	
	for range 4 {
		pen.Walk(passo)
	
		pen.Left(45)
		trigo(pen, dist * 0.3)
		
		pen.Right(90)
		
		trigo(pen, dist * 0.3)
		pen.Left(45)

	}

	pen.Up()
	pen.Walk(-dist)
	pen.Down()
	
} */
func main() {
	pen := NewPen(2000, 2000)
	pen.SetRGB(0, 0, 0)
	pen.SetHeading(0)
	pen.SetPosition(400, 400)
	pen.SetLineWidth(4)
	quadrados(pen, 1200)
	pen.SavePNG("tree.png")
	fmt.Println("PNG file created successfully.")
}
