package main

import "fmt"

func main() {
	pontoDeEbulicaoEmKelvin := 373
	var pontoDeEbulicaoEmCelsius = pontoDeEbulicaoEmKelvin - 273
	fmt.Printf("O ponto de ebulição em Kelvin é de %d, e o ponto de ebulição em Celsius é %d.", pontoDeEbulicaoEmKelvin, pontoDeEbulicaoEmCelsius)
}
