package TransportVersion2

import "fmt"

type Train struct {
	Destination string
}

func (t Train) PassengersIn() {
	fmt.Printf("Пассажири сіли у потяг до %sу \n", t.Destination)
}
func (t Train) PassengersOut() {
	fmt.Printf("Пассажири вийшли з потягу до %sу \n", t.Destination)
}
