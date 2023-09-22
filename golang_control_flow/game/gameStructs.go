package game

type Scene struct {
	Description string
	Choices     []Choice
}

type Choice struct {
	Text string
	Next *Scene
}
