package main

import "fmt"

func main() {
	var base, altura float64
	fmt.Printf("AREA DE UN TRIANGULO \n\n")
	fmt.Printf("Ingresa las medidas del triangulo...\n")
	fmt.Printf("base: ")
	fmt.Scanf("%f", &base)
	fmt.Printf("altura: ")
	fmt.Scanf("%f", &altura)
	area := (base * altura) / 2
	fmt.Println("\nArea: ", area)
}
