package TransportVersion2

type Transport interface {
	PassengersIn()
	PassengersOut()
}

func Passengers() {
	car := Car{"BMW"}
	train := Train{"Мілан"}
	airplane := Airplane{"Boeing 777"}

	vehicles := []Transport{car, train, airplane}

	for _, vehicle := range vehicles {
		vehicle.PassengersIn()
		vehicle.PassengersOut()
	}
}
