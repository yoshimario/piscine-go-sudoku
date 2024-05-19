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
	/* Use of 2D slice of bytes to represent the board instead of string increases memory usage and efficiency.
	Converting the empty spacese represesented by ',' into strings
	makess needing allocation of exra memory. Performance is also improved as bytes are faster with a lower level of abstraction.
	Then there is the character encoding in which GO uses the character encoding of UTF-8 characters which ASCII characters use mmore memory.
	Simplicity is also are more efficient as the characters are treated uniformy. And then if you wanted more flexibility to add even more rows it
	would be possible but then it be a whole new game.
	*/
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
			/* first part acceses the ith column of the rowth row of the board then ==num compares accessed element with number trying to be placed. if matched number
			exists in same row placement not valid. The colum works the same way only for the colth column and ithrow of the board. The 3x3 check  is the calcualation bit at the end
			since each game is a 3x3 box replicated 3 times across and three times down. The calculation (row/3)*3+i/3 determines the starting row index of the 3x3 box that includes the target position.
			(col/3)*3+i%3 calculates the starting column index of the 3x3 box that includes the target position. board[(row/3)*3+i/3][(col/3)*3+i%3]: This accesses the element at the calculated
			starting position of the 3x3 box. == num: Finally, this checks if the accessed element matches the number we're trying to place. If they match, it means the number already exists in the same 3x3 box,
			rendering the placement invalid. Use of modulus is needed to Adding i%3 adjusts the index based on the exact position of the cell within the 3x3 box. The % 3 operation ensures that the index wraps
			around correctly within the bounds of the 3x3 box.



			*/
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
