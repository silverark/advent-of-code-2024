package part1

import "log"

func process(input []string) int {
	var guard Guard
	guardStart := [2]int{0, 0}
	//Where is the guard
	found := false
	for row := range input {
		for col := range input[row] {
			if input[row][col] == '^' {
				guard = Guard{row: row, col: col, direction: 'U', inGrid: true, input: input}
				guard.visited = make(map[[3]int]bool)
				guard.visited[[3]int{row, col, 'U'}] = true
				found = true

				guardStart = [2]int{row, col}
				break
			}
		}
		if found {
			break
		}
	}

	originalGuard := guard
	loopCount := 0

	// Run once to see the guards path
	for {
		guard.move()
		if guard.inLoop == true {
			loopCount++
			break
		}
		if guard.inGrid == false {
			break
		}
	}

	visitedPath := make(map[[2]int]bool)
	for place := range guard.visited {
		visitedPath[[2]int{place[0], place[1]}] = true
	}

	for pathLocation := range visitedPath {
		row := pathLocation[0]
		col := pathLocation[1]

		guard = originalGuard
		guard.visited = make(map[[3]int]bool)
		guard.visited[[3]int{guard.row, guard.col, 'U'}] = true

		// ignore where guard started.
		if row == guardStart[0] && col == guardStart[1] {
			continue
		}
		original := string(input[row][col])

		// Put a # at teh row col position
		input[row] = input[row][:col] + "#" + input[row][col+1:]

		for {
			guard.move()

			if guard.inLoop == true {
				loopCount++
				break
			}

			if guard.inGrid == false {
				break
			}

			// Idiot protection
			if guard.steps > 100000 {
				log.Fatal("Guard has taken too many steps")
			}
		}

		// Put the original item back in to the string
		input[row] = input[row][:col] + original + input[row][col+1:]

	}

	return loopCount
}

type Guard struct {
	row       int
	col       int
	direction int
	inGrid    bool
	inLoop    bool
	visited   map[[3]int]bool
	input     []string
	steps     int
}

func (g *Guard) move() {
	nextSquare := g.nextSquare()
	//log.Printf("Next square: %v Guard Direction: %v. Col: %v Row: %v", nextSquare, g.direction, g.col, g.row)
	if nextSquare == "#" {
		g.turnRight()
		return
	}
	if g.inGrid == false || g.inLoop == true {
		return
	}
	g.moveForward()
	g.visited[[3]int{g.row, g.col, g.direction}] = true
}

func (g *Guard) nextSquare() string {
	newRow := g.row
	newCol := g.col
	if g.direction == 'U' {
		newRow--
	}
	if g.direction == 'D' {
		newRow++
	}
	if g.direction == 'L' {
		newCol--
	}
	if g.direction == 'R' {
		newCol++
	}
	// See if the guard is now out of bounds.
	if newRow < 0 || newRow >= len(g.input) || newCol < 0 || newCol >= len(g.input[newRow]) {
		g.inGrid = false
		return ""
	}
	// See if we're in a loop
	if g.visited[[3]int{newRow, newCol, g.direction}] == true {
		g.inLoop = true
		return ""
	}

	return string(g.input[newRow][newCol])
}

func (g *Guard) turnRight() {
	if g.direction == 'U' {
		g.direction = 'R'
		return
	}
	if g.direction == 'R' {
		g.direction = 'D'
		return
	}
	if g.direction == 'D' {
		g.direction = 'L'
		return
	}
	if g.direction == 'L' {
		g.direction = 'U'
		return
	}
}

func (g *Guard) moveForward() {
	if g.direction == 'U' {
		g.row--
	}
	if g.direction == 'D' {
		g.row++
	}
	if g.direction == 'L' {
		g.col--
	}
	if g.direction == 'R' {
		g.col++
	}
	g.steps++
}
