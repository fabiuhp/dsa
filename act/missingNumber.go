package main

import "fmt"

func main() {
	teste := []int{0, 2, 3}
	fmt.Println(missingNumber(teste))
}

func missingNumber(nums []int) int {
	tamanhoArray := len(nums)
	sum := tamanhoArray * (tamanhoArray + 1) / 2

	for _, valor := range nums {
		sum -= valor
	}

	return sum
}
