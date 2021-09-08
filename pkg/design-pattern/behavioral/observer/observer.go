package observer

import "fmt"

//go:generate mockery --all --keeptree

type Subjecter interface {
	Attach(o Observer)
	UpdateMsg(msg string)
}

var _ Subjecter = &Subject{}

type Subject struct {
	observers []Observer
}

func NewSubject() *Subject {
	return &Subject{
		observers: make([]Observer, 0),
	}
}

func (s *Subject) Attach(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *Subject) notify(msg string) {
	for _, o := range s.observers {
		o.Update(msg)
	}
}

func (s *Subject) UpdateMsg(msg string) {
	s.notify(msg)
}

type Observer interface {
	Update(msg string)
}

type Reader struct {
	name string
}

func NewReader(name string) *Reader {
	return &Reader{
		name: name,
	}
}

func (r *Reader) Update(msg string) {
	fmt.Printf("%s receive %s\n", r.name, msg)
}

func Happy(i string) string {
	return i
}
