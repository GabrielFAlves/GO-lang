package main

import "fmt"

func torreDeHanoi(n int, origem, auxiliar, destino string) {
	if n == 1 {
		fmt.Printf("Mova o disco 1 de %s para %s\n", origem, destino)
		return
	}

	torreDeHanoi(n-1, origem, destino, auxiliar)
	fmt.Printf("Mova o disco %d de %s para %s\n", n, origem, destino)
	torreDeHanoi(n-1, auxiliar, origem, destino)
}

func findPair(array []int, target int) (int, int) {
    left, right := 0, len(array)-1

    for left < right {
        currentSum := array[left] + array[right]

        if currentSum == target {
            return array[left], array[right]
        } else if currentSum < target {
            left++
        } else {
            right--
        }
    }

    // Nenhum par foi encontrado
    return -1, -1
}

func main() {
	//hanoi
	n := 3

	fmt.Printf("Solução para a Torre de Hanói com %d discos:\n", n)
	torreDeHanoi(n, "A", "B", "C")

	fmt.Println("********************************")
	fmt.Println("********************************")

    // problema 2
    array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    target := 10

    result1, result2 := findPair(array, target)

    if result1 == -1 && result2 == -1 {
        fmt.Println("Nenhum par encontrado.")
    } else {
        fmt.Printf("Par encontrado: (%d, %d)\n", result1, result2)
    }
}
