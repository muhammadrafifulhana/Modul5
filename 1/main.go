package main

import "fmt"

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {

	var suku int = 10

	fmt.Print("n")
	for i := 0; i <= suku; i++ {
		fmt.Print(" ", i)
	}

	fmt.Printf("\n")

	fmt.Print("Sn")
	for i := 0; i <= suku; i++ {
		fmt.Print(" ", fibonacci(i))
	}
}
