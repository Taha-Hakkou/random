package main

import (
	"os"

//	"fmt"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	var sudoku [9][9]rune
	for i := 0; i < len(args); i++ {
		for j := 0; j < len(args[i]); j++ {
			sudoku[i][j] = rune(args[i][j])
		}
	}
	printSudoku(sudoku)
	z01.PrintRune('\n')
	var x = [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _ = range x {
		z01.PrintRune('\n')
		solveSudoku(&sudoku)
	}
	printSudoku(sudoku)
	//}
}

//func isValidSudoku() bool {

//}

func solveSudoku(sudoku *[9][9]rune) {
	var values = [3]int{1, 2, 4}
	var indexes = [8]int{-1, 0, 1, -1, 2, -1, -1, -1}
	for x := '1'; x <= '9'; x++ {
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				a := 0
				b := 0
				found := false
				for m := 0; m < 3; m++ {
					for n := 0; n < 3; n++ {
						if sudoku[3*i+m][3*j+n] == x {
							found = true
							break
						} else if sudoku[3*i+m][3*j+n] == '.' {
							a += values[m]
							//
							notAdded := false
							if b == 0 || b == 7-values[n]{
								notAdded = true
							}
							for k := 0; k < 3 && k != n; k++ {
								if b == values[k] {
									notAdded = true
								}
							}
							if notAdded {
								b += values[n]
							}
							break
						}
					}
					if found {
						break
					}
				}
				if found {
					continue
				}
				for k := 0; k < 3 && k != j; k++ {
					for m := 0; m < 3; m++ {
						n := 0
						for n = 0; n < 3; n++ {
							if sudoku[3*i+m][3*k+n] == x {
								a -= values[m]
								break
							}
						}
						if n != 3 {
							break
						}
					}
				}
				for k := 0; k < 3 && k != i; k++ {
					for m := 0; m < 3; m++ {
						n := 0
						for n = 0; n < 3; n++ {
							if sudoku[3*k+m][3*j+n] == x {
								b -= values[n]
								break
							}
						}
						if n != 3 {
							break
						}
					}
				}
				if indexes[a] != -1 && indexes[b] != -1 {
					sudoku[3*i+indexes[a]][3*j+indexes[b]] = x
				}
			}
		}
	}
}

func printSudoku(sudoku [9][9]rune) {
	for i := 0; i < len(sudoku); i++ {
		z01.PrintRune(sudoku[i][0])
		for j := 1; j < len(sudoku[i]); j++ {
			z01.PrintRune(' ')
			z01.PrintRune(sudoku[i][j])
		}
		z01.PrintRune('\n')
	}
}
