package main

import (
	"fmt"
	"math"
	"math/rand"
	"projeto2/geometria"
)

type Ponto struct {
	X float64
	Y float64
}

func main() {

	printArray()

	//****************************************************

	var inputString string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&inputString)
	inverted := reverseString(inputString)
	fmt.Println("String invertida:", inverted)

	//****************************************************

	ponto := Ponto{X: 9, Y: 25}
	distancia := ponto.DistanciaOrigem()
	fmt.Printf("Distância até a origem: %.2f\n", distancia)

	//****************************************************

	var largura, altura float64
    fmt.Print("Digite a largura do retângulo: ")
    fmt.Scanln(&largura)
    fmt.Print("Digite a altura do retângulo: ")
    fmt.Scanln(&altura)

    area := geometria.CalcularArea(largura, altura)
    perimetro := geometria.CalcularPerimetro(largura, altura)

    fmt.Println("Área do retângulo:", area)
    fmt.Println("Perímetro do retângulo:", perimetro)
}

func (p *Ponto) DistanciaOrigem() float64 {
	distancia := math.Sqrt(p.X*p.X + p.Y*p.Y)
	return distancia
}

func reverseString(str string) string {

	runes := []rune(str)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func printArray() {

	var array [10]int

	for i := 0; i < 10; i++ {
		array[i] = rand.Intn(100)
	}

	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}
}
