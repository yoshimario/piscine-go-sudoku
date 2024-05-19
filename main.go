package main

import (
	"fmt"
	"os"
)

const gridSize = 9 // The size of the Sudoku grid is 9x9

func main() {
	// Check if the number of arguments is correct
	if len(os.Args) != gridSize+1 {
		fmt.Println("Error")
		return
	}

	// Create the Sudoku board from input
	board := make([][]byte, gridSize)
	for i := range board {
		// Each argument should be 9 characters long
		if len(os.Args[i+1]) != gridSize {
			fmt.Println("Error")
			return
		}
		board[i] = []byte(os.Args[i+1])
	}

	// Check if the initial board is valid
	if !isBoardValid(board) {
		fmt.Println("Error")
		return
	}

	// Try to solve the Sudoku
	if solveSudoku(board) {
		printBoard(board) // Print the solved board
	} else {
		fmt.Println("Error") // Print error if no solution
	}
}

// Check if placing 'num' at board[row][col] is valid
func isValid(board [][]byte, row, col int, num byte) bool {
	for i := 0; i < gridSize; i++ {
		// Check the row, column, and 3x3 box
		if board[row][i] == num || board[i][col] == num || board[(row/3)*3+i/3][(col/3)*3+i%3] == num {
			return false
		}
	}
	return true
}

// Solve the Sudoku using backtracking
func solveSudoku(board [][]byte) bool {
	for row := 0; row < gridSize; row++ {
		for col := 0; col < gridSize; col++ {
			if board[row][col] == '.' { // Find an empty cell
				for num := byte('1'); num <= '9'; num++ {
					if isValid(board, row, col, num) { // Try placing numbers 1-9
						board[row][col] = num
						if solveSudoku(board) { // Recursively solve the rest
							return true
						}
						board[row][col] = '.' // Undo if not successful
					}
				}
				return false // Return false if no number works
			}
		}
	}
	return true // Return true if solved
}

// Print the board
func printBoard(board [][]byte) {
	for row := range board {
		for col := range board[row] {
			fmt.Printf("%c", board[row][col])
			if col < gridSize-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

// Check if the initial board is valid
func isBoardValid(board [][]byte) bool {
    for row := 0; row < gridSize; row++ {
        for col := 0; col < gridSize; col++ {
            if board[row][col] != '.' {
                num := board[row][col]
                if num < '1' || num > '9' { // Check if the character is between '1' and '9'
                    return false
                }
                board[row][col] = '.'
                if !isValid(board, row, col, num) {
                    return false
                }
                board[row][col] = num
            }
        }
    }
    return true
}
