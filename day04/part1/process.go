package part1

func process(input []string) int {
	lineCnt := len(input)
	total := 0
	for row := 0; row < lineCnt; row++ {
		for col := 0; col < len(input[row]); col++ {
			if input[row][col] == 'X' {
				found := findWord(input, row, col)
				total += found
			}
		}
	}
	return total
}

// Look through all the directions of the crossword to find the word XMAS from this starting letter which will be an X
func findWord(input []string, row, col int) int {

	findCount := 0
	// Check for XMAS right
	if col+3 < len(input[row]) && input[row][col+1] == 'M' && input[row][col+2] == 'A' && input[row][col+3] == 'S' {
		findCount++
	}
	// left direction
	if col-3 >= 0 && input[row][col-1] == 'M' && input[row][col-2] == 'A' && input[row][col-3] == 'S' {
		findCount++
	}
	// down
	if row+3 < len(input) && input[row+1][col] == 'M' && input[row+2][col] == 'A' && input[row+3][col] == 'S' {
		findCount++
	}
	// up
	if row-3 >= 0 && input[row-1][col] == 'M' && input[row-2][col] == 'A' && input[row-3][col] == 'S' {
		findCount++
	}
	// diagonal down right
	if row+3 < len(input) && col+3 <= len(input[row]) && input[row+1][col+1] == 'M' && input[row+2][col+2] == 'A' && input[row+3][col+3] == 'S' {
		findCount++
	}
	// diagonal up right
	if row-3 >= 0 && col+3 <= len(input[row]) && input[row-1][col+1] == 'M' && input[row-2][col+2] == 'A' && input[row-3][col+3] == 'S' {
		findCount++
	}
	// diag down left
	if row+3 < len(input) && col-3 >= 0 && input[row+1][col-1] == 'M' && input[row+2][col-2] == 'A' && input[row+3][col-3] == 'S' {
		findCount++
		//log.Printf("Found XMAS diagonal down left: %v", string(input[row][col])+string(input[row+1][col-1])+string(input[row+2][col-2])+string(input[row+3][col-3]))
	}
	// diag up left
	if row-3 >= 0 && col-3 >= 0 && input[row-1][col-1] == 'M' && input[row-2][col-2] == 'A' && input[row-3][col-3] == 'S' {
		findCount++
	}
	return findCount
}
