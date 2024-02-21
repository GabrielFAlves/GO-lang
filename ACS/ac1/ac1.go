package main

import "fmt"

func calculaMedia (numbers ...float64) float64 {
	soma := 0.0
	for i := 0; i < len(numbers); i++ {
        soma += numbers[i]
    }

    media := soma / float64(len(numbers))
    return media
}

func verificaParidade (num int) string {
	if num % 2 == 0 {
		return "par"
	} else {
		return "ímpar"
	}
}

func minhaPotencia (base int, expoente int) int {
	resultado := 1

	for i := 0; i < expoente; i++ {
		resultado *= base
	}

	return resultado
}

func converteCelsiusParaFahrenheit (celsius float64) float64 {
	return (celsius * 9/5) + 32
}

func main () {
	fmt.Println(calculaMedia(5,6,8,5))
	fmt.Println(verificaParidade(5))
	fmt.Println(minhaPotencia(3,3))
	fmt.Println(converteCelsiusParaFahrenheit(40))
}