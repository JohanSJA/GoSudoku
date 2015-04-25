// Reference: https://freepythontips.wordpress.com/2013/09/01/sudoku-solver-in-python/
package main

import (
	"fmt"
	"os"
	"strings"
)

func sameRow(i, j int) bool {
	return i/9 == j/9
}

func sameCol(i, j int) bool {
	return (i-j)%9 == 0
}

func sameBlock(i, j int) bool {
	return i/27 == j/27 && i%9/3 == j%9/3
}

func r(a string) {
	i := strings.Index(a, "0")
	if i == -1 {
		fmt.Println(a)
		os.Exit(0)
	}

	excludedNumbers := make(map[string]bool)
	for j := 0; j < 81; j++ {
		if sameRow(i, j) || sameCol(i, j) || sameBlock(i, j) {
			excludedNumbers[string(a[j])] = true
		}
	}

	for _, m := range "123456789" {
		if _, ok := excludedNumbers[string(m)]; !ok {
			r(a[:i] + string(m) + a[i+1:])
		}
	}
}

func main() {
	if len(os.Args) != 2 || len(os.Args[1]) != 81 {
		fmt.Println("Usage: sudoku puzzle")
		fmt.Println("  where puzzle is an 81 character string representing the puzzle read left-to-right, top-to-bottom, and 0 is a blank")
		os.Exit(1)
	}
	r(os.Args[1])
}
