package TransportTypes

import "fmt"

type PassengerTrain struct {
	Destination string
}

func (t PassengerTrain) PassengersIn() {
	fmt.Printf("Пассажири сіли у потяг до %sу \n", t.Destination)
}
func (t PassengerTrain) PassengersOut() {
	fmt.Printf("Пассажири вийшли з потягу до %sу \n", t.Destination)
}
