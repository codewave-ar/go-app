package main

import (
	"fmt"
	"go-app-01/operaciones"
)

func main() {
	const num1 int = 2
	const num2 int = 4
	simbolo := "/"

	defer fmt.Println("Fin de la funcion")
	defer fmt.Println("Se hizo la operacion ", num1, simbolo, num2)

	switch simbolo {
	case "+":
		resultadoSuma := operaciones.Suma(num1, num2)
		fmt.Println("El resultado de la suma es: ", resultadoSuma)
	case "/":
		resultadoDivision, err := operaciones.Division(num1, num2)
		if err != nil {
			fmt.Println("Error intentando dividir: ", err.Error())
		} else {
			fmt.Println("El resultado de la division es: ", resultadoDivision)
		}
	case "^":
		resultadoPotencia := operaciones.Potencia(num1, num2)
		fmt.Println("El resultado de la potencia es: ", resultadoPotencia)
	case "-":
		fallthrough
	default:
		fmt.Println("La operacion que intenta hacer, no esta disponible")
	}
}
