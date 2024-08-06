package main

import (
	"piscine"
	"fmt"
)

func printNbr(n int) {
	fmt.Print(n)
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	piscine.ForEach(printNbr, a)
}
