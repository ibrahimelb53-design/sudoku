package main

import  "fmt"

func main() {
	slice := []string{".96.4...1","1...6...4","5.481.39.","..795..43",".3..8....","4.5.23.18",".1.63..59",".59.7.83.","..359...7"}
	board := intgrid(slice)

if solveSudoku(board){
	for _, row := range board {
    fmt.Println(row)
}
}else{
	fmt.Println("no solution")
}


}
func intgrid(input []string) [][]int {
    grid := [][]int{}

    for _, line := range input {
        row := []int{}
        for _, ch := range line {
            if ch == '.' {
                row = append(row, 0)
            } else {
                row = append(row, int(ch-'0'))
            }
        }
        grid = append(grid, row)
    }
    return grid
}


func solveSudoku(board [][]int) bool{
	for i := 0 ; i < 9 ; i++{
		for j := 0 ; j < 9 ; j++{
			row := i
			col := j
			if board[row][col] == 0 {
				for c := 1 ; c <= 9 ; c++{
					if isValid(board, col, row, c){
						board[row][col] = c
						if solveSudoku(board){
							return true
						}
						board[row][col] = 0
					}
				}
				return false
			}
		}
	}
	return true
}

func isValid(board[][]int,col int,row int , c int) bool{
	for i := 0 ; i < 9 ; i++{
		if board[row][i] == c {
			return false
		}
		if board[i][col] == c {
			return false
		}
		boxrow := 3*(row/3) + i/3
		boxcol := 3*(col/3) + i%3
		if board[boxrow][boxcol] == c {
			return false
		}
	}
	return true
}
	// 3 9 6   2 4 5   7 8 1  	096  040  001
	// 1 7 8   3 6 9   5 2 4	100  060  004
	// 5 2 4   8 1 7   3 9 6	504  810  390
	
	// 2 8 7   9 5 1   6 4 3	007  950  043
	// 9 3 1   4 8 6   2 7 5	030  080  000
	// 4 6 5   7 2 3   9 1 8	405  023  018
	
	// 7 1 2   6 3 8   4 5 9	010  630  059
	// 6 5 9   1 7 4   8 3 2	059  070  830
	// 8 4 3   5 9 2   1 6 7	003  590  007
	//

