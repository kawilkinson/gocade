package snake

type Snake struct {
	Body      []Coordinate
	Length    int
	Direction int
}

type Coordinate struct {
	x, y int
}

type Food struct {
	x, y int
}

func (s *Snake) GetSnakeHead() Coordinate {
	return s.Body[len(s.Body)-1]
}

func (s *Snake) HitWall(m *SnakeGameModel) bool {
	head := s.GetSnakeHead()

	if head.x >= m.Height || head.y >= m.Width-1 || head.x <= 0 || head.y <= 0 {
		return true
	}

	return false
}

func (s *Snake) HitSelf(coord Coordinate) bool {
	for _, snakePart := range s.Body {
		if snakePart.x == coord.x && snakePart.y == coord.y {
			return true
		}
	}

	return false
}
