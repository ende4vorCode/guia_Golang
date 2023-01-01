// La diferencia entre variable y constante es que la constante no puede cambiar su valor

package main

import "fmt"

func main() {

	// Declaración de constantes
	const pi float64 = 3.141516
	const pi2 = 3.14

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	// Declaración de variables enteras
	base := 12 // Cuando la variable no ha sido declarada previamente, podemos usar :=
	var altura int = 14
	var area float64 = float64(base) * float64(altura) // Cuando utilizamos flotantes en una variable, GO nos cambia el tipo de dato de los datos que estamos usando. Esto puede llevar a un problema de memoria, hay que tener cuidado.

	fmt.Println("El area es:", area)

	// Zero values: Son variables declaradas, pero no tienen un valor.
	var a int     // 0
	var b float32 // 0
	var c float64 // 0
	var d string  // void
	var e bool    // false

	fmt.Println(a, b, c, d, e)
}
