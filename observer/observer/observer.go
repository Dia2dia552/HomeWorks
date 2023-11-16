package observer

import "fmt"

type Observer interface {
	Update(string)
}
type Subject struct {
	observers []Observer
}

type Player struct {
	Name string
}

func (s *Subject) Register(o Observer) {
	s.observers = append(s.observers, o)
}
func (s *Subject) Deregister(o Observer) {
	for i, observer := range s.observers {
		if observer == o {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}
func (s *Subject) Notify(msg string) {
	for _, observer := range s.observers {
		observer.Update(msg)
	}
}

func (p *Player) Update(msg string) {
	fmt.Printf("[%s] Received update: %s\n", p.Name, msg)
}
