package main

import "fmt"

func printBintang(n int) {
	if n > 0 {
		printBintang(n - 1)
		fmt.Println(bintang(n))
	}
}

func bintang(n int) string {
	if n == 1 {
		return "*"
	}
	return "*" + bintang(n-1)
}

func main() {
	var n int
	fmt.Scan(&n)

	printBintang(n)
}
