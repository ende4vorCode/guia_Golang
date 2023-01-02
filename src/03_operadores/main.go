package main

import "fmt"

func main() {
	const x = 10
	const y = 5
	var result float64

	// Suma
	result = x + y // 15
	fmt.Println("El resultado de la suma es:", result)

	// Resta
	result = x - y // 5
	fmt.Println("El resultado de la resta es:", result)

	// Multiplicación

	result = x * y // 50
	fmt.Println("El resultado de la multiplicación es:", result)

	// Division

	result = x / y // 2
	fmt.Println("El resultado de la division es:", result)

	// Modulo

	result = x % y
	fmt.Println("El modulo de la operación es:", result)

	// Incremental
	incremental := 0
	fmt.Println("El valor inicial de incremental es:", incremental)
	incremental++
	fmt.Println("El valor final de incremental es:", incremental)

	// Decremental
	incremental--
	fmt.Println("El nuevo valor de incremental es:", incremental)

}
