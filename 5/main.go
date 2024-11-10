package main

import "fmt"

func printGanjil(n, current int) {
	if current > n {
		return
	}
	if current%2 != 0 {
		fmt.Printf("%d ", current)
	}
	printGanjil(n, current+1)
}

func main() {
	var n int
	fmt.Scan(&n)

	printGanjil(n, 1)
}
