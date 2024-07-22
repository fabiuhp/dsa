package main

import "time"

// Exemplo simples de Time Complexity
func main() {
	timeStart := time.Now()
	findSum(9999999)
	// findSumIterate(9999999)
	timeEnd := time.Now()

	println(timeEnd.Sub(timeStart).Milliseconds())
}

func findSum(n int) int {
	return n * (n + 1) / 2
}

func findSumIterate(n int) (sum int) {
	for i := 0; i < n; i++ {
		sum += 1
	}
	return
}
