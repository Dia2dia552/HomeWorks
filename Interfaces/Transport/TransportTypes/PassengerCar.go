package TransportTypes

import "fmt"

type PassengerCar struct {
	Name string
}

func (c PassengerCar) PassengersIn() {
	fmt.Printf("Пассажири сіли у авто %s\n", c.Name)
}
func (c PassengerCar) PassengersOut() {
	fmt.Printf("Пассажири вийшли з авто %s\n", c.Name)
}
