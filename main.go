package main

import  "fmt"

func main() {
	slice := []string{".96.4...1","1...6...4","5.481.39.","..795..43",".3..8....","4.5.23.18",".1.63..59",".59.7.83.","..359...7"}
	board := switcher(slice)
	if solveSudoku(board) {
	for _, row := range board {
		fmt.Println(row)
	}
} else {
	fmt.Println("no result")
}

}

func switcher(slice []string) [][]int {
    board := [][]int{}

    for i := 0; i < len(slice); i++ {
        row := []int{}
        for j := 0; j < len(slice[i]); j++ {
            if slice[i][j] == '.' {
                row = append(row, 0)
            } else {
                row = append(row, int(slice[i][j]-'0'))
            }
        }
        board = append(board, row)
    }
    return board
}

	func solveSudoku(board [][]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				for c := 1; c <= 9; c++ {
					if isValid(board, i, j, c) {
						board[i][j] = c
						if solveSudoku(board) {
							return true
						}
						board[i][j] = 0
					}
				}
				return false
			}
		}
	}
	return true
}
	



	// --- BOX ---  //

	// box
	// 1 2 3
	// 4 5 6
	// 7 8 9
	

	// 3 9 6   2 4 5   7 8 1  	096  040  001
	// 1 7 8   3 6 9   5 2 4	100  060  004
	// 5 2 4   8 1 7   3 9 6	504  810  390
	
	// 2 8 7   9 5 1   6 4 3	007  950  043
	// 9 3 1   4 8 6   2 7 5	030  080  000
	// 4 6 5   7 2 3   9 1 8	405  023  018
	
	// 7 1 2   6 3 8   4 5 9	010  630  059
	// 6 5 9   1 7 4   8 3 2	059  070  830
	// 8 4 3   5 9 2   1 6 7	003  590  007

 


	



func make3x3() [][]int {
    matrix := make([][]int, 3)
    for i := range 3 {
        matrix[i] = []int{}
    }
    return matrix
}

func isValid(board [][]int, row, col, c int) bool {

	
	for i := 0; i < 9; i++ {
		if board[row][i] == c {
			return false
		}
		if board[i][col] == c {
			return false
		}
	}

for i := 0; i < 9; i++ {
	for j := 0; j < 9; j++ {

		
if i-3 < 0 {
    if j-3 < 0 {
        if row-3 < 0 && col-3 < 0 {
            if board[i][j] == c {
                return false
            }
        }
    } else if j-3 < 3 {
        if row-3 < 0 && col-3 >= 0 && col-3 < 3 {
            if board[i][j] == c {
                return false
            }
        }
    } else {
        if row-3 < 0 && col-3 >= 3 {
            if board[i][j] == c {
                return false
            }
        }
    }
} else if i-3 < 3 {
    if j-3 < 0 {
        if row-3 >= 0 && row-3 < 3 && col-3 < 0 {
            if board[i][j] == c {
                return false
            }
        }
    } else if j-3 < 3 {
        if row-3 >= 0 && row-3 < 3 && col-3 >= 0 && col-3 < 3 {
            if board[i][j] == c {
                return false
            }
        }
    } else {
        if row-3 >= 0 && row-3 < 3 && col-3 >= 3 {
            if board[i][j] == c {
                return false
            }
        }
    }
} else {
    if j-3 < 0 {
        if row-3 >= 3 && col-3 < 0 {
            if board[i][j] == c {
                return false
            }
        }
    } else if j-3 < 3 {
        if row-3 >= 3 && col-3 >= 0 && col-3 < 3 {
            if board[i][j] == c {
                return false
            }
        }
    } else {
        if row-3 >= 3 && col-3 >= 3 {
            if board[i][j] == c {
                return false
            }
        }
    }
}

		}
	}
	return true
}


//        0 1 2 | 3 4 5 | 6 7 8
      
 // 0 |   1 1 1 | 2 2 2 | 3 3 3
  //1 |   1 1 1 | 2 2 2 | 3 3 3
  //2 |   1 1 1 | 2 2 2 | 3 3 3
      
  //3 |   4 4 4 | 5 5 5 | 6 6 6
  //4 |   4 4 4 | 5 5 5 | 6 6 6
  //5 |   4 4 4 | 5 5 5 | 6 6 6
      
  //6 |   7 7 7 | 8 8 8 | 9 9 9
  //7 |   7 7 7 | 8 8 8 | 9 9 9
  //8 |   7 7 7 | 8 8 8 | 9 9 9









