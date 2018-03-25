package game

type group struct {
	Colour    Colour
	Positions []int
	Liberties int
}

func (g group) Contains(i int) bool {
	// TODO this lookup needs to be fast
	// maybe we could order positions
	// or lookup by index
	for p := 0; p < len(g.Positions); p++ {
		if g.Positions[p] == i {
			return true
		}
	}
	return false
}

func findGroup(stones []Colour, boardSize, x, y int) group {
	if !insideBoard(boardSize, x, y) {
		return group{Colour: None}
	}
	i := boardIndex(boardSize, x, y)
	col := stones[i]
	if col == None {
		return group{Colour: None}
	}
	g := group{Colour: col, Positions: []int{i}}
	g = walkGroup(stones, g, boardSize, x, y)
	return g
}

func walkGroup(stones []Colour, g group, boardSize, x, y int) group {
	neighbours := getNeighbours(x, y)
	for _, n := range neighbours {
		if !insideBoard(boardSize, n.x, n.y) {
			continue
		}

		ni := boardIndex(boardSize, n.x, n.y)
		if g.Contains(ni) {
			continue // We've already walked this cell
		}

		switch stones[ni] {
		case None:
			g.Liberties++
		case g.Colour:
			g.Positions = append(g.Positions, ni)
			g = walkGroup(stones, g, boardSize, n.x, n.y)
		default:
			// This is an opponents stone
			// Don't increment liberties
			// and don't walk the group
			// any further
		}
	}

	return g
}

func insideBoard(boardSize, x, y int) bool {
	return x >= 0 && y >= 0 && x < boardSize && y < boardSize
}

func boardIndex(boardSize, x, y int) int {
	return x + y*boardSize
}

type pos struct{ x, y int }

func getNeighbours(x, y int) [4]pos {
	return [4]pos{
		{x - 1, y},
		{x, y - 1},
		{x + 1, y},
		{x, y + 1},
	}
}
