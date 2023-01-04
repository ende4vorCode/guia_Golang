// Los programas en Go siempre deben tener la extension .go y comienzan con un archivo main.go

package main // Un package es la carpeta donde esta guardado, en este caso al ser el archivo principal, lo llamamos main

import "fmt" // Si no definimos las funciones que vamos a importar, la extension de GO lo hace por su cuenta, asi como la eliminación de las librerías que no ocupemos.

func main() { // Creamos la función principal
	fmt.Println("Hola mundo! Este es mi primer programa en GO!") // fmt es la función que nos permite imprimir por consola y Println() es la función que permite escribir un mensaje.
}

// Antes de ejecutar cualquier programa, debemos hacer un build de este. go build src/01_holaMundo.go
// AL hacerlo, veremos que nos crea un archivo el cual es el archivo que compilamos.

// También podemos utilizar go run src/01_holaMundo.go para ejecutarlo sin tener que compilar. Esto es menos eficiente, pero para el desarrollo es mas cómodo.
