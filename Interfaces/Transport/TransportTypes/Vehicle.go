package TransportTypes

import "fmt"

type Transport interface {
	Stop()
	Move()
	ChangeSpeed(speed int)
}
type PassengerTransport interface {
	AddPassengers()
	RemovePassengers()
}

func Moving() {
	car := Car{Name: "Автомобіль", Speed: 60}
	train := Train{Name: "Потяг", Speed: 120}
	airplane := Airplane{Name: "Літак", Speed: 800, Altitude: 10000}

	vehicles := []Transport{car, train, airplane}

	for _, vehicle := range vehicles {
		vehicle.Move()
		vehicle.ChangeSpeed(+10)
		vehicle.Stop()
	}
}

func Passengers() {
	car := PassengerCar{"BMW", 2}
	train := PassengerTrain{"Мілан", 50}
	airplane := PassengerAirplane{"Boeing 777", 100}

	vehicles := []PassengerTransport{car, train, airplane}

	for _, vehicle := range vehicles {
		vehicle.AddPassengers()
		vehicle.RemovePassengers()
	}
}

func SeeTheRoute() {
	car := Car{Name: "Автомобіль", Speed: 60}
	train := Train{Name: "Потяг", Speed: 120}
	airplane := Airplane{Name: "Літак", Speed: 800, Altitude: 10000}

	route := NewRoute()
	route.AddTransport(car)
	route.AddTransport(train)
	route.AddTransport(airplane)

	fmt.Println("Список транспортних засобів на маршруті:")

	for _, transport := range route.GetTransportList() {
		fmt.Printf("Тип: %T\n", transport)
	}
}
