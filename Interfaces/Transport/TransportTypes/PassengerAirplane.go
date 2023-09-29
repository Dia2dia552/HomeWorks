package TransportTypes

import (
	"fmt"
)

type PassengerAirplane struct {
	Name       string
	Passengers int
}

func (a PassengerAirplane) AddPassengers() {
	a.Passengers++
	fmt.Printf("Пассажири сіли у літак %s\n", a.Name)
}
func (a PassengerAirplane) RemovePassengers() {
	if a.Passengers > 0 {
		a.Passengers--
		fmt.Printf("Пасажири вийшли з літаку %s\n", a.Name)
	} else {
		fmt.Println("В літаку більше немає пасажирів")
	}

}
