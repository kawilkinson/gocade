package snake

type Model struct {
	HorizontalLine string
	VerticalLine   string
	EmptySymbol    string
	SnakeSymbol    string
	FoodSymbol     string
	Stage          [][]string
	Snake          Snake
	GameOver       bool
	Score          int
	Food           Food

	Width          int
	Height         int
}
