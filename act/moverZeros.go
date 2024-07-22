package main

import "fmt"

func main() {
	fmt.Println(moverZeros([]int{0, 1, 0, 3, 12}))
}

func moverZeros(array []int) []int {
	indexNaoZero := 0

	for i := 0; i < len(array); i++ {
		if array[i] != 0 {
			array[indexNaoZero] = array[i]
			indexNaoZero++
		}
	}

	for i := indexNaoZero; i < len(array); i++ {
		array[i] = 0
	}

	return array
}
