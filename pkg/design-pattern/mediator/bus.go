package mediator

import "fmt"

type Bus interface {
	Arrive()
	Depart()
	PermitArrival()
}

type IntercityBus struct {
	Mediator mediator
}

var _ Bus = &IntercityBus{}

func (ib *IntercityBus) Arrive() {
	if !ib.Mediator.canArrive(ib) {
		fmt.Println("Intercity Bus: Arrival blocked, waiting....")
		return
	}
	fmt.Println("Intercity: Arrived!!!")
}

func (ib *IntercityBus) Depart() {
	ib.Mediator.notifyAboutDeparture()
}
func (ib *IntercityBus) PermitArrival() {
	fmt.Println("IntercityBus: Arrival permitted")
	ib.Arrive()
}
type ShuttleBus struct {
	Mediator mediator
}


func (sb *ShuttleBus) Arrive() {
	if !sb.Mediator.canArrive(sb) {
		fmt.Println("ShuttleBus Bus: Arrival blocked, waiting....")
		return
	}
	fmt.Println("Intercity: Arrived!!!")
}

func (sb *ShuttleBus) Depart() {
	sb.Mediator.notifyAboutDeparture()
}
func (sb *ShuttleBus) PermitArrival() {
	fmt.Println("ShuttleBus: Arrival permitted")
	sb.Arrive()
}
