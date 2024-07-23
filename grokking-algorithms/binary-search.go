package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := BinarySearch([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 13)
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
	} else {
		fmt.Printf("Valor final: %v\n", result)
	}
}

func BinarySearch(array []int, chave int) (int, error) {
	menor := 0
	maior := len(array) - 1

	for menor <= maior {
		meio := (menor + maior) / 2
		posicao := array[meio]

		if posicao == chave {
			return posicao, nil
		} else if posicao > chave {
			maior = meio - 1
		} else {
			menor = meio + 1
		}
	}
	return 0, errors.New("Não encontrado.")
}
