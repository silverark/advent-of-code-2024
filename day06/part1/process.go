package part1

func process(input []string) int {
	var guard Guard
	//Where is the guard
	found := false
	for row := range input {
		for col := range input[row] {
			if input[row][col] == '^' {
				guard = Guard{row: row, col: col, direction: "U", inGrid: true, input: input}
				guard.visited = make(map[[2]int]bool)
				guard.visited[[2]int{row, col}] = true
				found = true
				break
			}
		}
		if found {
			break
		}
	}
	for {
		guard.move()
		if guard.inGrid == false {
			break
		}
	}
	return len(guard.visited)
}

type Guard struct {
	row       int
	col       int
	direction string
	inGrid    bool
	visited   map[[2]int]bool
	input     []string
	steps     int
}

func (g *Guard) move() {
	nextSquare := g.nextSquare()
	if nextSquare == "#" {
		g.turnRight()
		return
	}
	if g.inGrid == false {
		return
	}
	g.moveForward()
	g.visited[[2]int{g.row, g.col}] = true
}

func (g *Guard) nextSquare() string {
	newRow := g.row
	newCol := g.col
	if g.direction == "U" {
		newRow--
	}
	if g.direction == "D" {
		newRow++
	}
	if g.direction == "L" {
		newCol--
	}
	if g.direction == "R" {
		newCol++
	}
	// See if the guard is now out of bounds.
	if newRow < 0 || newRow >= len(g.input) || newCol < 0 || newCol >= len(g.input[newRow]) {
		g.inGrid = false
		return ""
	}
	return string(g.input[newRow][newCol])
}

func (g *Guard) turnRight() {
	if g.direction == "U" {
		g.direction = "R"
		return
	}
	if g.direction == "R" {
		g.direction = "D"
		return
	}
	if g.direction == "D" {
		g.direction = "L"
		return
	}
	if g.direction == "L" {
		g.direction = "U"
		return
	}
}

func (g *Guard) moveForward() {
	if g.direction == "U" {
		g.row--
	}
	if g.direction == "D" {
		g.row++
	}
	if g.direction == "L" {
		g.col--
	}
	if g.direction == "R" {
		g.col++
	}
	g.steps++
}
