package main

import "fmt"

// Una función sirve para evitar que repitamos código haciéndolo modular y fácil de leer

func greetings() { // Función saludo
	var message = "Hola mundo!"
	fmt.Println(message)
}

func sum(a, b int) { // Función que suma dos argumentos, es una buena practica declararlo como esta aquí
	fmt.Println("sum:", a+b)
}

func sayHi(name string) {
	fmt.Println("HI", name)
}

func returnValue(a, b int) int { // a,b int son los argumentos y el int de fuera es el valor del string
	return a + b
}

func doubleOf(a int) (b, c int) { // Función que nos permite regresar 2 valores distintos.
	return a, a * 2
}

func main() { // La función main es el punto de entrada de nuestro programa
	greetings()                                    // Hola mundo!
	sum(5, 4)                                      //9
	sayHi("Ende4vor")                              // HI Ende4vor
	fmt.Println("returnValue:", returnValue(5, 4)) //9 Cual es la diferencia entre la func sum y returnValue? sum solamente imprime un valor, pero ese valor no lo podemos utilizar posteriormente, en cambio returnValue si regresa ese valor.
	var single, double int = doubleOf(5)           // Asignamos 2 valores
	fmt.Println("doubleOf:", single, double)       // 5 10
	_, double = doubleOf(3)                        // Esta forma de asignar valores, noes permite elegir cual de los dos utilizar
	fmt.Println(double)                            // 6

}
