package funcs

import (
	"math"
)

var mySquare [][]string

func Solve(tetrominosArray [][4][4]string) [][]string {
	// Initialize the board with a starting dimension of 4x4.
	// If we can't place all tetrominoes, we'll increase the size by 1 and initialize the board again.
	size := int(math.Ceil(math.Sqrt(float64(4 * len(tetrominosArray)))))

	// Initialize the square grid.
	mySquare = InitSquare(size)

	// Continue to solve using backtracking until a solution is found.
	for !BacktrackSolver(tetrominosArray, 0) {
		// Increase the grid size by 1.
		size++
		// Reinitialize the square grid with the new size.
		mySquare = InitSquare(size)
	}

	// Return the filled square grid as a 2D string array.
	return mySquare
}

func InitSquare(n int) [][]string {
	//initializes a square
	var Square [][]string
	var row []string
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			row = append(row, ".")
		}
		Square = append(Square, row)
		row = []string{}
	}
	return Square
}

func BacktrackSolver(tetrominoes [][4][4]string, n int) bool {
	if n == len(tetrominoes) { //base condition when all tetrominoes are placed, board is solved
		return true
	}

	for i := 0; i < len(mySquare); i++ {
		for j := 0; j < len(mySquare); j++ {
			if CheckInsert(i, j, tetrominoes[n]) { //check if we can place current tetrominoe on the board anywhere
				Insert(i, j, tetrominoes[n]) // if we can place it at this location, check if we can place another piece
				if BacktrackSolver(tetrominoes, n+1) {
					return true
				}
				Remove(i, j, tetrominoes[n]) //if the next piece can't be placed, backtrack
			}
		}
	} // if we can't place tetro anywhere, return false
	return false
}
func CheckInsert(i, j int, tetro [4][4]string) bool { //check if we can place piece at current location
	for a := 0; a < 4; a++ {
		for b := 0; b < 4; b++ {
			if tetro[a][b] != "." {
				if i+a == len(mySquare) || j+b == len(mySquare) || mySquare[i+a][j+b] != "." {
					return false
				}
			}
		}

	}
	return true
}

func Insert(i, j int, tetro [4][4]string) { // insert piece and when all 4 pieces "#" are placed, no need to place '.'
	a, b, c := 0, 0, 0
	for a < 4 {
		for b < 4 {
			if tetro[a][b] != "." {
				c++
				mySquare[i+a][j+b] = tetro[a][b]
				if c == 4 {
					break
				}
			}
			b++
		}
		b = 0
		a++
	}
}

func Remove(i, j int, tetro [4][4]string) { //remove piece at current location
	a, b, c := 0, 0, 0
	for a < 4 {
		for b < 4 {
			if tetro[a][b] != "." {
				if c == 4 {
					break
				}
				mySquare[i+a][j+b] = "."
			}
			b++
		}
		b = 0
		a++
	}
}
