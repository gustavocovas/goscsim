package goscsim

type Event struct {
	Time  int
	Actor Actor
}

type EventQueue interface {
	Pop() *Event
	Push(e *Event)
	Len() int
}
