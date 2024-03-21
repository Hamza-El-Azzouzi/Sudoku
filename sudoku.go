package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

const gridSize = 9

func printGrid(grid [][]int) {
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			z01.PrintRune(rune(grid[i][j] + '0'))
			z01.PrintRune(' ')
		}
		fmt.Println()
	}
}

func findEmptyLocation(grid [][]int) (int, int, bool) {
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			if grid[i][j] == 0 {
				return i, j, true
			}
		}
	}
	return -1, -1, false
}

func isSafe(grid [][]int, row, col, num int) bool {
	return isValidGrid(grid) && !usedInRow(grid, row, num) &&
		!usedInCol(grid, col, num) &&
		!usedInBox(grid, row-row%3, col-col%3, num)
}

func usedInRow(grid [][]int, row, num int) bool {
	for i := 0; i < gridSize; i++ {
		if grid[row][i] == num {
			return true
		}
	}
	return false
}

func usedInCol(grid [][]int, col, num int) bool {
	for i := 0; i < gridSize; i++ {
		if grid[i][col] == num {
			return true
		}
	}
	return false
}

func usedInBox(grid [][]int, boxStartRow, boxStartCol, num int) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[i+boxStartRow][j+boxStartCol] == num {
				return true
			}
		}
	}
	return false
}

func solveSudoku(grid [][]int) bool {
	row, col, found := findEmptyLocation(grid)
	if !found {
		return true
	}

	for num := 1; num <= 9; num++ {
		if isSafe(grid, row, col, num) {
			grid[row][col] = num
			if solveSudoku(grid) {
				return true
			}
			grid[row][col] = 0
		}
	}

	return false
}
func isValidGrid(grid [][]int) bool {
	for i := 0; i < gridSize; i++ {
		if !isValidRow(grid, i) || !isValidCol(grid, i) || !isValidBox(grid, i-i%3, i%3*3) {
			return false
		}
	}
	return true
}

func isValidRow(grid [][]int, row int) bool {
	seen := make(map[int]bool)
	for _, num := range grid[row] {
		if num != 0 {
			if seen[num] {
				return false
			}
			seen[num] = true
		}
	}
	return true
}

func isValidCol(grid [][]int, col int) bool {
	seen := make(map[int]bool)
	for i := 0; i < gridSize; i++ {
		num := grid[i][col]
		if num != 0 {
			if seen[num] {
				return false
			}
			seen[num] = true
		}
	}
	return true
}

func isValidBox(grid [][]int, startRow, startCol int) bool {
	seen := make(map[int]bool)
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			num := grid[i][j]
			if num != 0 {
				if seen[num] {
					return false 
				}
				seen[num] = true
			}
		}
	}
	return true
}

func main() {
	if len(os.Args) <= 9 {
		fmt.Println("Error.")
		return
	}

	grid := make([][]int, gridSize)
	for i := range grid {
		grid[i] = make([]int, gridSize)
	}

	for i := 1; i < len(os.Args); i++ {
		for j, char := range os.Args[i] {
			if char != '.' {
				num, err := strconv.Atoi(string(char))
				if err != nil || num < 1 || num > 9 {
					fmt.Println("Error")
					return
				}
				grid[i-1][j] = num
			}
		}
	}

	if solveSudoku(grid) {
		printGrid(grid)

	} else {
		fmt.Println("Error")
	}
}
