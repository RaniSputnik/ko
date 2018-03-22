package game

import "log"

func recalculateLiberties(stones []Colour, boardSize int, x, y int) (result []Colour, captured int) {
	log.Printf("recalculateLiberties::boardSize=%d,x=%d,y=%d", boardSize, x, y)
	result = stones

	i := boardIndex(boardSize, x, y)
	col := stones[i]
	if col == None {
		return
	}
	var captureCol Colour
	if col == Black {
		captureCol = White
	} else {
		captureCol = Black
	}

	neighbours := getNeighbours(x, y)
	for _, n := range neighbours {
		if !hasAtLeastOneLiberty(result, boardSize, n.x, n.y) {
			result, captured = captureRegion(result, captureCol, boardSize, n.x, n.y)
		}
	}

	return
}

func hasAtLeastOneLiberty(stones []Colour, boardSize int, x, y int) bool {
	log.Printf("hasAtLeastOneLiberty::boardSize=%d,x=%d,y=%d", boardSize, x, y)
	liberyFound, noLibertyFound := true, false

	// TODO does this need to live here?
	// we probably already know that x/y is within the board boundaries
	if !insideBoard(boardSize, x, y) {
		return liberyFound // TODO should this be an error
	}

	i := boardIndex(boardSize, x, y)
	col := stones[i]
	if col == None {
		return liberyFound
	}

	neighbours := getNeighbours(x, y)
	for _, n := range neighbours {
		if !insideBoard(boardSize, n.x, n.y) {
			continue
		}

		i := n.x + n.y*boardSize
		if ncol := stones[i]; ncol == None {
			return liberyFound
		} else if ncol != col {
			continue // No liberty here
		}

		if hasAtLeastOneLiberty(stones, boardSize, n.x, n.y) {
			return liberyFound
		}
	}

	return noLibertyFound
}

func captureRegion(stones []Colour, colour Colour, boardSize int, x, y int) (result []Colour, captured int) {
	log.Printf("captureRegion::colour=%s,boardSize=%d,x=%d,y=%d", colour, boardSize, x, y)
	result = stones

	if i := boardIndex(boardSize, x, y); result[i] != None {
		result[i] = None
		captured++
	}

	neighbours := getNeighbours(x, y)
	for _, n := range neighbours {
		if !insideBoard(boardSize, n.x, n.y) {
			continue
		}

		i := boardIndex(boardSize, n.x, n.y)
		if result[i] == colour {
			result[i] = None
			captured++
		}
	}
	return
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
