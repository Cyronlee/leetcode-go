package main

import (
	"fmt"
)

/*
037 解数独 https://leetcode-cn.com/problems/sudoku-solver

编写一个程序，通过填充空格来解决数独问题。

一个数独的解法需遵循如下规则：

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
空白格用 '.' 表示。
*/
func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	printBoard(board)
	fmt.Println("---------------------------")
	solveSudoku(board)
	printBoard(board)
}

var row, col [9][9]bool
var block [3][3][9]bool
var queue [][2]int

// 执行用时  ms 内存消耗  MB
func solveSudoku(board [][]byte) {
	// 必须初始化变量才能过leetcode
	row, col, block, queue = [9][9]bool{}, [9][9]bool{}, [3][3][9]bool{}, [][2]int{}
	for x, line := range board {
		for y, b := range line {
			if b == '.' {
				queue = append(queue, [2]int{x, y})
			} else {
				number := b - '1'
				row[x][number] = true
				col[y][number] = true
				block[x/3][y/3][number] = true
			}
		}
	}
	dfs(0, board)
}

// 深度优先算法(递归+回溯)
func dfs(pos int, board [][]byte) bool {
	if len(queue) == pos {
		return true
	}
	x, y := queue[pos][0], queue[pos][1]
	for number := byte(0); number < 9; number++ {
		if !(row[x][number] || col[y][number] || block[x/3][y/3][number]) {
			row[x][number], col[y][number], block[x/3][y/3][number] = true, true, true
			board[x][y] = number + '1'
			if dfs(pos+1, board) {
				return true
			}
			row[x][number], col[y][number], block[x/3][y/3][number] = false, false, false
		}
	}
	return false
}

func printBoard(board [][]byte) {
	for _, row := range board {
		for _, n := range row {
			fmt.Printf(" %v ", string(n))
		}
		fmt.Println()
	}
}
