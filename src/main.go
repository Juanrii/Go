package main

import "fmt"

func calcularArea(base, altura int) {
	result := base * altura
	fmt.Println(result)
}

func main() {
	// Constantes
	const pi float64 = 3.14
	const pi2 = 3.14
	fmt.Println(pi, pi2)

	// Variables enteras
	base := 12 // := declara y asgina
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a,b,c,d)

	// Sprintf
	name := "Juan"
	message := fmt.Sprintf("Hi, I'm %s", name)
	fmt.Println(message)

	// Tipo de dato %T
	fmt.Printf("name: %T", name)

	calcularArea(2,4)
}	