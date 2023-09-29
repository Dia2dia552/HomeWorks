package TransportTypes

import "fmt"

type PassengerCar struct {
	Name       string
	Passengers int
}

func (c PassengerCar) AddPassengers() {
	c.Passengers++
	fmt.Printf("Пассажири сіли у авто %s\n", c.Name)
}
func (c PassengerCar) RemovePassengers() {
	if c.Passengers > 0 {
		c.Passengers--
		fmt.Printf("Пасажири вийшли з авто %s\n", c.Name)
	} else {
		fmt.Println("В машині нема більше пасажирів")
	}

}
