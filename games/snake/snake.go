package snake

type Snake struct {
	body      []Coordinate
	length    int
	direction int
}

type Coordinate struct {
	x, y int
}

type Food struct {
	x, y int
}

func (s *Snake) getSnakeHead() Coordinate {
	return s.body[len(s.body)-1]
}

func (s *Snake) hitWall(m *Model) bool {
	head := s.getSnakeHead()

	if head.x >= m.Height || head.y >= m.Width-1 || head.x <= 0 || head.y <= 0 {
		return true
	}

	return false
}

func (s *Snake) hitSelf(coord Coordinate) bool {
	for _, snakePart := range s.body {
		if snakePart.x == coord.x && snakePart.y == coord.y {
			return true
		}
	}

	return false
}
