package TransportVersion1

import "fmt"

type Airplane struct {
	Name     string
	Speed    int
	Altitude int
}

func (a Airplane) Move() {
	fmt.Printf("%s летить зі швидкістю %d км/год на висоті %d метрів\n", a.Name, a.Speed, a.Altitude)

}
func (a Airplane) Stop() {
	fmt.Printf("%s приземлився\n", a.Name)
}

func (a Airplane) ChangeSpeed(speed int) {
	a.Speed = speed
	fmt.Printf("%s змінює швидкість на %d км/год\n", a.Name, a.Speed)
}
