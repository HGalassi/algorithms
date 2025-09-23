package main

import "fmt"

func Fatorial(n int) int {
	if n <= 0 {
		return 1
	}
	return n * Fatorial(n-1)
}

func Challenge(n int) int {
	if n == 1 {
		fmt.Print(" Challenge -> Exiba os números de 1 a 10 recursivamente ")
	}

	if n > 10 {
		return 0
	}

	fmt.Print(n, " ")
	return Challenge(n + 1)
}
