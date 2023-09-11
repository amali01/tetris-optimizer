package funcs

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func ReadInput(file io.Reader) [][4][4]string {
	// Initialize variables to store tetromino data and parsing state
	var tetrominoArray [][4][4]string
	var tetromino [4][4]string
	scanner := bufio.NewScanner(file)
	lineCount, tetroNum, newLineFlag := 0, 0, true

	// Read lines from the input file
	for scanner.Scan() {
		line := scanner.Text()

		// Check for an empty line indicating the end of a tetromino block
		if line == "" {
			// Handle empty line based on parsing state
			if newLineFlag && !(tetroNum == 0) {
				newLineFlag = false
				continue
			} else {
				fmt.Printf("ERROR: at line %d\n", lineCount+1+(tetroNum*5))
				os.Exit(0)
			}
		} else if lineCount == 0 && tetroNum != 0 && line != "" && newLineFlag {
			// Detect errors if line count, tetromino count, and empty line conditions are not met
			fmt.Printf("ERROR: at line %d\n", lineCount+(tetroNum*5))
			os.Exit(0)
		}

		// Process each character in the line to construct the tetromino
		var arr [4]string
		if len(line) != 4 {
			// Check for invalid line length
			fmt.Printf("ERROR: at line %d\n", lineCount+1+(tetroNum*5))
			os.Exit(0)
		}
		for ind := range arr {
			if rune(line[ind]) == '.' {
				arr[ind] = "."
			} else if rune(line[ind]) == '#' {
				arr[ind] = string(rune('A' + tetroNum))
			} else {
				// Handle errors for invalid characters
				fmt.Printf("ERROR: at line %d\n", lineCount+1+(tetroNum*5))
				os.Exit(0)
			}
		}

		// Store the constructed tetromino and update parsing state
		tetromino[lineCount] = arr
		lineCount++
		if lineCount == 4 {
			newLineFlag = true
			lineCount = 0
			tetroNum++

			// Check for invalid tetromino format and optimize the tetromino
			if !CheckTetromino(tetromino) {
				fmt.Printf("ERROR: invalid tetromino format at tetroNum %d\n", tetroNum)
				os.Exit(0)
			}
			tetromino = OptimizeTetromino(tetromino)
			tetrominoArray = append(tetrominoArray, tetromino)
		}
	}

	// Check for scanner errors and handle them
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return tetrominoArray
}

func OptimizeTetromino(tetromino [4][4]string) [4][4]string {
	// Optimize the given tetromino by shifting empty rows and columns
	// fmt.Println("tetromino before", tetromino)

	// Shift empty rows to the top
	i := 0
	for {
		zeroes := 0
		for j := 0; j < 4; j++ {
			if tetromino[i][j] == "." {
				zeroes++
			}
		}
		if zeroes == 4 { // If the entire row is empty, shift it by 1 row to the top
			tetromino = ShiftVertical(tetromino)
			continue
		}
		break
	}

	// Shift empty columns to the left
	for {
		zeroes := 0
		for j := 0; j < 4; j++ {
			if tetromino[j][i] == "." {
				zeroes++
			}
		}
		if zeroes == 4 { // If the entire column is empty, shift it by 1 column to the left
			tetromino = ShiftHorizontal(tetromino)
			continue
		}
		break
	}

	// fmt.Println("tetromino after", tetromino)

	return tetromino
}

func ShiftVertical(tetromino [4][4]string) [4][4]string {
	// Shifts the given tetromino one row upwards

	// Store the first row in a temporary variable
	temp := tetromino[0]
	// Shift rows up by one position
	tetromino[0] = tetromino[1]
	tetromino[1] = tetromino[2]
	tetromino[2] = tetromino[3]
	tetromino[3] = temp
	return tetromino
}

func ShiftHorizontal(tetromino [4][4]string) [4][4]string {
	// Shifts the given tetromino one column to the left

	// Transpose the tetromino to treat columns as rows
	tetromino = Transpose(tetromino)
	// Shift rows (which were columns after transposition) up by one position
	tetromino = ShiftVertical(tetromino)
	// Transpose the tetromino back to its original orientation
	tetromino = Transpose(tetromino)
	return tetromino
}

func Transpose(tetromino [4][4]string) [4][4]string {
	// Transposes the given tetromino by swapping rows and columns

	// Create a temporary tetromino to store the transposed values
	temp := tetromino
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			// Swap the values between rows and columns
			temp[i][j] = tetromino[j][i]
		}
	}
	return temp
}

func CheckTetromino(tetromino [4][4]string) bool {
	// Initialize counters for adjacent blocks (conactedSideCoun) and total block count (blockCount)
	conactedSideCoun, blockCount := 0, 0

	// Iterate through each row (a) and column (b) in the 4x4 tetromino array
	for x, row := range tetromino {
		for y, cell := range row {
			// Check if the current cell is not empty (contains x block)
			if cell != "." {
				// Increment the total block count
				blockCount++

				// Check adjacent cells in all four directions (up, down, left, right)
				// and increment conactedSideCoun for each adjacent block
				if x+1 < 4 && tetromino[x+1][y] != "." {
					conactedSideCoun++
				}
				if x-1 >= 0 && tetromino[x-1][y] != "." {
					conactedSideCoun++
				}
				if y+1 < 4 && tetromino[x][y+1] != "." {
					conactedSideCoun++
				}
				if y-1 >= 0 && tetromino[x][y-1] != "." {
					conactedSideCoun++
				}
			}
		}
	}

	// Check conditions for x valid tetromino:
	// 1. The total block count must be 4.
	if blockCount != 4 {
		return false
	}

	// 2. The number of adjacent blocks around each block must be 6 or 8.
	if conactedSideCoun == 6 || conactedSideCoun == 8 {
		return true
	}

	// If neither condition is met, return false
	return false
}
