package main

import "fmt"

func printFaktor(n, i int) {
	if i > n {
		return
	}
	if n%i == 0 {
		fmt.Printf("%d ", i)
	}
	printFaktor(n, i+1)
}

func main() {
	var n int
	fmt.Scan(&n)

	printFaktor(n, 1)
}
