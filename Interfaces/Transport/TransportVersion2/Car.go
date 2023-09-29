package TransportVersion2

import "fmt"

type Car struct {
	Name string
}

func (c Car) PassengersIn() {
	fmt.Printf("Пассажири сіли у авто %s\n", c.Name)
}
func (c Car) PassengersOut() {
	fmt.Printf("Пассажири вийшли з авто %s\n", c.Name)
}
