package main

import "fmt"

// fmt es el paquete que utilizamos para interactuar por consola, pero ademas tiene otras funcionalidades

func main() {
	// Declaración de variables
	helloMessage := "Hello"
	worldMessage := "World"

	// Println - Es una función que permite escribir por consola y al final le agrega una linea
	fmt.Println("Println:", helloMessage)
	fmt.Println("Println:", worldMessage)

	// Printf - Es una función que permite escribir por consola y agrega una función extra al valor que le indiques como paramento

	name := "Ende4vor"
	age := 25
	// %s indica que va a recibir un string y %d que va a recibir un entero
	// \n indica un salto de linea
	// Posteriormente indica los valores que va a recibir la impresión
	fmt.Printf("Printf: %s tiene %d años \n", name, age) // Buena practica.
	// En caso de no saber que tipo de dato va, pondremos %v
	fmt.Printf("Printf: %v tiene %v años \n", name, age)

	// Sprintf - Genera un string pero no lo muestra por consola, solo lo guarda como tal. Entonces se puede usar en variables.

	message := fmt.Sprintf("Sprint: %s tiene %d años", name, age)
	fmt.Println(message)

	// Tipo de datos - Con el paquete fmt podemos ademas identificar tipos de datos.
	fmt.Printf("helloMessage: %T \n", helloMessage)
	fmt.Printf("age: %T \n", age)

}

// Para mas detalle https://pkg.go.dev/fmt
