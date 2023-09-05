package arrays

func isValidSudoku(board [][]byte) bool {

	var rows, columns, squares [9][9]bool

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			current_val := board[r][c]

			if current_val == '.' {
				continue
			}

			index := int(current_val-'0') - 1

			if rows[r][index] || columns[c][index] || squares[r/3*3+c/3][index] {
				return false
			}
			rows[r][index], columns[c][index], squares[r/3*3+c/3][index] = true, true, true
		}
	}

	return true
}
