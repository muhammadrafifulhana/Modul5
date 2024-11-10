package main

import "fmt"

func printDesc(n int) {
	if n < 1 {
		return
	}
	fmt.Printf("%d ", n)
	printDesc(n - 1)
}

func printAsc(n, curr int) {
	if curr > n {
		return
	}
	fmt.Printf("%d ", curr)
	printAsc(n, curr+1)
}

func main() {
	var n int
	fmt.Scan(&n)

	printDesc(n)
	printAsc(n, 1)
}
