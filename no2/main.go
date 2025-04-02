package main

import "fmt"

// Faktorial menggunakan rekursi
func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	num := 5
	fmt.Printf("Faktorial dari %d adalah %d\n", num, factorial(num))
}
