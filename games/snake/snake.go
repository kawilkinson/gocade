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

	if head.x >= m.Height-1 || head.y >= m.Width-1 || head.x < 1 || head.y < 1 {
		return true
	}

	return false
}

// Extra check to prevent panics from going out of bounds, to be fixed later when I have time
func ExtraHitWallCheck(m *SnakeGameModel, coord Coordinate) bool {
	return coord.x >= m.Height-1 || coord.y >= m.Width-1 || coord.x < 1 || coord.y < 1
}

func (s *Snake) HitSelf(coord Coordinate) bool {
	for _, snakePart := range s.Body {
		if snakePart.x == coord.x && snakePart.y == coord.y {
			return true
		}
	}

	return false
}
