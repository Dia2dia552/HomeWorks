package TransportTypes

type Route struct {
	TransportList []Transport
}

func NewRoute() *Route {
	return &Route{}
}

func (r *Route) AddTransport(transport Transport) {
	r.TransportList = append(r.TransportList, transport)
}

func (r *Route) GetTransportList() []Transport {
	return r.TransportList
}
