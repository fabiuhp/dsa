package main

func missingNumber(nums []int) int {
	tamanhoArray := len(nums)
	sum := tamanhoArray * (tamanhoArray + 1) / 2

	for _, valor := range nums {
		sum -= valor
	}

	return sum
}
