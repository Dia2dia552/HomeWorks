package TransportTypes

import (
	"fmt"
)

type PassengerAirplane struct {
	Name string
}

func (a PassengerAirplane) PassengersIn() {
	fmt.Printf("Пассажири сіли у літак %s\n", a.Name)
}
func (a PassengerAirplane) PassengersOut() {
	fmt.Printf("Пассажири вийшли з літаку %s\n", a.Name)
}
