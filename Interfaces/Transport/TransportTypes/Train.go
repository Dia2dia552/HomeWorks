package TransportTypes

import "fmt"

type Train struct {
	Name  string
	Speed int
}

func (t Train) Move() {
	fmt.Printf("%s рухається зі швидкістю %d км/год \n", t.Name, t.Speed)
}

func (t Train) Stop() {
	fmt.Printf("%s зупиняється\n", t.Name)
}

func (t Train) ChangeSpeed(speed int) {
	t.Speed += speed
	fmt.Printf("%s змінює швидкість на %d км/год\n", t.Name, t.Speed)

}
