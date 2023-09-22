package TransportVersion1

type Transport interface {
	Stop()
	Move()
	ChangeSpeed(speed int)
}

func Moving() {
	car := Car{Name: "Автомобіль", Speed: 60, Engine: "ДВС"}
	train := Train{Name: "Потяг", Speed: 120}
	airplane := Airplane{Name: "Літак", Speed: 800, Altitude: 10000}

	vehicles := []Transport{&car, train, airplane}

	for _, vehicle := range vehicles {
		vehicle.Move()
		//vehicle.ChangeSpeed(vehicle.Speed + 10)
		vehicle.Stop()
	}
}
