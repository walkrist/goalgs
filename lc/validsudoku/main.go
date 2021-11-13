package main

import (
	"fmt"
)

func initMap(filMap map[int]int) map[int]int {
	for n := 1; n <= 9; n++ {
		filMap[n] = 0
	}
	return filMap
}

func isValidSudoku(board [][]int) bool {
	myMap := make(map[int]int)
	for row := 0; row < 9; row++ {
		myMap = initMap(myMap)
		for i := 0; i < 9; i++ {
			if myMap[(board[row][i])] > 1 && board[row][i] > 0 && board[row][i] <= 9 {
				return false
			}
			myMap[int(board[row][i])]++
		}
	}
	for col := 0; col < 9; col++ {
		myMap = initMap(myMap)
		for i := 0; i < 9; i++ {
			if myMap[int(board[i][col])] > 1 && board[i][col] > 0 && board[i][col] <= 9 {
				return false
			}
			myMap[int(board[i][col])]++
		}
	}
	return true
}

func main() {
	fmt.Println(isValidSudoku([][]int{[]int{5, 3, 0, 0, 7, 0, 0, 0, 0}, []int{6, 0, 0, 1, 9, 5, 0, 0, 0}, []int{0, 9, 8, 0, 0, 0, 0, 6, 0}, []int{8, 0, 0, 0, 6, 0, 0, 0, 3}, []int{4, 0, 0, 8, 0, 3, 0, 0, 1}, []int{7, 0, 0, 0, 2, 0, 0, 0, 6}, []int{0, 6, 0, 0, 0, 0, 2, 8, 0}, []int{0, 0, 0, 4, 1, 9, 0, 0, 5}, []int{0, 0, 0, 0, 8, 0, 0, 7, 9}}))
}
