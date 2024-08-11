package main

import "fmt"

func main() {
	fmt.Println(isPalindrome("arara"))
}

func isPalindrome(word string) bool {
	start := 0
	end := len(word) - 1

	for start < end {
		primeraLetra := word[start]
		ultimaLetra := word[end]

		if primeraLetra != ultimaLetra {
			return false
		}

		start++
		end--
	}

	return true
}
