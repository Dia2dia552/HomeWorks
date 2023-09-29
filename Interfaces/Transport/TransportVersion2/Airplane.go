package TransportVersion2

import (
	"fmt"
)

type Airplane struct {
	Name string
}

func (a Airplane) PassengersIn() {
	fmt.Printf("Пассажири сіли у літак %s\n", a.Name)
}
func (a Airplane) PassengersOut() {
	fmt.Printf("Пассажири вийшли з літаку %s\n", a.Name)
}
