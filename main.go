package main

import (
	"AMJ/funcs"
	"fmt"
	"os"
	"time"
)

// var mu sync.Mutex
var mySquare [][]string

func main() {
	if len(os.Args) != 2 {
		fmt.Println("[USAGE]: go run . sample.txt")
		return
	}
	fileName := os.Args[1]
	if fileName != "" {
		startTime := time.Now() // Record the start time

		file, err := os.Open(fileName)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}
		defer func() {
			if err = file.Close(); err != nil {
				fmt.Println(err)
				os.Exit(0)
			}
		}()
		tetrominosArray := funcs.ReadInput(file)

		mySquare = funcs.Solve(tetrominosArray)

		// Calculate and print the execution time
		elapsedTime := time.Since(startTime)
		fmt.Println("...........................................")
		fmt.Printf("Execution time: %v\n", elapsedTime)

		// Print the solution
		PrintSolution(mySquare, fileName)
	}
}

func PrintSolution(mySquare [][]string, fileName string) {
	fmt.Printf("Testing file:( %s )\n\n", fileName)
	for i := range mySquare {
		for j := range mySquare {
			fmt.Printf("%s ", mySquare[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Println("...........................................")

}
