package main

import "fmt"

func main() {
	valores := inverteArray([]int{1, 2, 0, 3, 4})
	fmt.Println(valores)
}

func inverteArray(arr []int) []int {
	inicio := 0
	final := len(arr) - 1

	for inicio < final {
		arr[inicio], arr[final] = arr[final], arr[inicio]
		inicio++
		final--
	}

	return arr
}
