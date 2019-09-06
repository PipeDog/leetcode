/*
 * @lc app=leetcode.cn id=36 lang=golang
 *
 * [36] 有效的数独
 */
func isValidSudoku(board [][]byte) bool {

	// enum row
	for col := 0; col < 9; col++ {
		eachRow := map[byte]bool{}

		for row := 0; row < 9; row++ {
			ch := board[col][row]

			if _, ok := eachRow[ch]; ok && ch != '.' {
				return false
			} else {
				eachRow[ch] = true
			}
		}
	}

	// enum col
	for row := 0; row < 9; row++ {
		eachCol := map[byte]bool{}

		for col := 0; col < 9; col++ {
			ch := board[col][row]

			if _, ok := eachCol[ch]; ok && ch != '.' {
				return false
			} else {
				eachCol[ch] = true
			}
		}
	}

	// enum 3*3
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {

			dict := map[byte]bool{}

			for row := i; row < i + 3; row++ {
				for col := j; col < j + 3; col++ {
					ch := board[row][col]

					if _, ok := dict[ch]; ok && ch != '.' {
						return false
					} else {
						dict[ch] = true
					}
				}
			}

		}
	}

	return true
}

