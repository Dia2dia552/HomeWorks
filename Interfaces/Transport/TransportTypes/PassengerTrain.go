package TransportTypes

import "fmt"

type PassengerTrain struct {
	Destination string
	Passengers  int
}

func (t PassengerTrain) AddPassengers() {
	t.Passengers++
	fmt.Printf("Пассажири сіли у потяг до %sу \n", t.Destination)
}
func (t PassengerTrain) RemovePassengers() {
	if t.Passengers > 0 {
		t.Passengers--
		fmt.Printf("Пасажири вийшли з потягу до %sу \n", t.Destination)
	} else {
		fmt.Println("В потязі нема пасажирів")
	}

}
