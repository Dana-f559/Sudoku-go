package main

import "fmt"

func main() {
	board := [][]int{ 
        {-1, -1, -1, -1, 8, 2, 5, -1, 1}, 
        {-1, 6, -1, -1, 5, -1, -1, -1, -1}, 
        {2, -1, -1, -1, -1, 9, -1, 3, 7}, 

        {-1, 9, 6, -1, -1, -1, -1, 5, 8},
		{-1, -1, -1, -1, -1, 7, 3, -1, -1}, 
        {5, -1, 3, -1, 9, -1, -1, -1, -1},

        {-1, 8, -1, 1, 7, -1, -1, -1, -1}, 
        {6, 7, 2, 3, 4, -1, 9, -1, 5},
		{1, -1, 5, 9, 2, -1, 8, 7, 4}, 

    } 

	result := solve_sudoku(board)

    if result {
        fmt.Println("Sudoku puzzle solved successfully!")
		fmt.Println(board)
    } else {
        fmt.Println("Could not solve Sudoku puzzle.")
    }
}


func find_next_empty(puzzle [][]int) (int, int) {
    for r := 0; r < 9; r++ {
        for c := 0; c < 9; c++ {
            if puzzle[r][c] == -1 {
                return r, c
            }
        }
    }
    return -2, -2 // If no empty cell is found
}


func is_valid(puzzle [][]int, guess int, row int, col int) bool{

	row_vals := puzzle[row]

	for i := 0; i< 9; i++ {
		if row_vals[i] == guess{
			return false
		}
	}

	var col_vals []int
    for i := 0; i < len(puzzle); i++ {
        col_vals = append(col_vals, puzzle[i][col])
    }
   
	for i := 0; i< 9; i++ {
		if col_vals[i] == guess{
			return false
		}
	}

	rowStart := (row / 3) * 3
	colStart := (col / 3) * 3

	// Checks if the guess is in the 3 by 3 grid
	for r := rowStart; r < rowStart+3; r++ {
		for c := colStart; c < colStart+3; c++ {
			// If yes, then the guess is False
			if puzzle[r][c] == guess {
				return false
			}
		}
	}

	return true
}
func solve_sudoku(puzzle [][]int) bool{


	row, col := find_next_empty(puzzle)

	if row == -2 {
		return true
	}

	for guess := 1; guess < 11; guess++ {
		if is_valid(puzzle, guess, row, col) {
			puzzle[row][col] = guess

			if solve_sudoku(puzzle) {
				// fmt.Println(puzzle)
				return true
			}
		}

		puzzle[row][col] = -1
	}

	return false
}
