package part1

func process(input []string) int {
	lineCnt := len(input)
	total := 0
	for row := 0; row < lineCnt; row++ {
		for col := 0; col < len(input[row]); col++ {
			if input[row][col] == 'A' {
				found := findWord(input, row, col)
				total += found
			}
		}
	}
	return total
}

// Look through all the directions of the crossword to find the word MAS from this starting letter which will be an A. Should only be diagonal
func findWord(input []string, row, col int) int {
	// Make sure teh A is at least 1 from each side.
	if row == 0 || col == 0 || row+1 >= len(input) || col+1 >= len(input[row]) {
		return 0
	}

	if input[row-1][col-1] == 'M' && input[row+1][col+1] == 'S' || input[row-1][col-1] == 'S' && input[row+1][col+1] == 'M' {
		// find /
		if input[row-1][col+1] == 'M' && input[row+1][col-1] == 'S' || input[row-1][col+1] == 'S' && input[row+1][col-1] == 'M' {
			return 1
		}
	}
	return 0
}
