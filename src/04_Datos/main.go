package main

import "fmt"

// GO tiene distintos de tipos de datos y al utilizarlos con un tipado fuerte, se puede ganar mucho rendimiento en las aplicaciones.

func main() {

	// Enteros
	var entero int     // En caso de no seleccionar el tipo de dato, GO seleccionara en base a tu sistema operativo 32 - 64
	var entero8 int8   // 8 bits = -1278 a 127
	var entero16 int16 // 16 bits = -2^15 a (2^15)-1
	var entero32 int32 // 32 bits = -2^31 a (2^31)-1
	var entero64 int64 // 63 bits = -2^63 a (2^63)-1

	// Enteros positivos
	var enteroPositivo uint     // OS default
	var enteroPositivo8 uint8   // 8 bits = 0 a (2^8)-1
	var enteroPositivo16 uint16 // 16 bits = 0 a (2^16)-1
	var enteroPositivo32 uint32 // 32 bits = 0 a (2^32)-1
	var enteroPositivo64 uint64 // 64 bits = 0 a (2^64)-1

	// Decimales

	var flotante32 float32 // 32 bits = +/- 1.18e^38 a +/- 3.4e^38
	var flotante64 float64 // 64 bits = +/- 2.23e^-308 a +/- 1.8e^308

	// Texto

	var texto string = "" // ""

	// Bool

	var booleano bool // true o false

	// Números complejos

	var complejo64 complex64   // Real e imaginario float32
	var complejo128 complex128 // Real o imaginario float64

	// ejemplo

	c := 10 + 8i
	fmt.Println(c) // 10 + 8i

	fmt.Println(entero, entero8, entero16, entero32, entero64, enteroPositivo, flotante32, flotante64, enteroPositivo, enteroPositivo16, enteroPositivo8, enteroPositivo32, enteroPositivo64, texto, booleano, complejo64, complejo128) // Esto es para evitar las rayas rojas xd
}
