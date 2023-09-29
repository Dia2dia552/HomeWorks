package TransportTypes

import "fmt"

type Car struct {
	Name  string
	Speed int
}

func (c Car) Move() {
	fmt.Printf("%s рухається зі швидкістю %d км/год\n", c.Name, c.Speed)
}

func (c Car) Stop() {
	fmt.Printf("%s зупиняється\n", c.Name)
}

func (c Car) ChangeSpeed(speed int) {
	c.Speed += speed
	fmt.Printf("%s змінює швидкість на %d км/год\n", c.Name, c.Speed)
}
