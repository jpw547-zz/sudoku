package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	board := [9][9]int{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			var num int
			counter := 0
			for {
				num = rand.Intn(9) + 1
				if !usedInRow(board, i, num) && !usedInCol(board, j, num) {
					break
				}
				counter++
				if counter > 9 {
					break
				}
			}
			board[i][j] = num
		}
	}

	printBoard(board)

}

func printBoard(board [9][9]int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%d ", board[i][j])
			if (j+1)%3 == 0 {
				fmt.Print("  ")
			}
		}
		fmt.Println("")
		if i != 8 && (i+1)%3 == 0 {
			fmt.Println("")
		}
	}
}

func usedInRow(board [9][9]int, rowIndex int, checkNum int) bool {
	for j := 0; j < 9; j++ {
		if board[rowIndex][j] == checkNum {
			return true
		}
	}
	return false
}

func usedInCol(board [9][9]int, colIndex int, checkNum int) bool {
	for i := 0; i < 9; i++ {
		if board[i][colIndex] == checkNum {
			return true
		}
	}
	return false
}
