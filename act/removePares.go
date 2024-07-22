package main

func main() {
	array := []int{1, 2, 3, 4}
	retorno := removeNumerosPares(array[:])

	for i := 0; i < len(retorno); i++ {
		println(retorno[i])
	}
}

func removeNumerosPares(array []int) (numerosPares []int) {
	for i := 0; i < len(array); i++ {
		if array[i]%2 == 0 {
			numerosPares = append(numerosPares, array[i])
		}
	}
	return
}
