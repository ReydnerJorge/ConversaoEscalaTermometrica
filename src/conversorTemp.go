package main

import "fmt"

// Funcao para converter de Kelvin para Celsius

func kelvinToCelsius(kelvin float64) float64 {
	return kelvin - 273.15
}

func main() {
	var kelvin float64

	// Solicita oo usuário a temperatura em Kelvin
	fmt.Print("Digite a temperatura em Kelvin: ")
	fmt.Scan(&kelvin)

	// Converte a temperatura para Celsius
	celsius := kelvinToCelsius(kelvin)

	// Exite o resultado
	fmt.Printf("A temperatura em Celsius é: %.2f°C\n", celsius)
}
