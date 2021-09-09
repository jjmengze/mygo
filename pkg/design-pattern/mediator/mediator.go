package mediator

type mediator interface {
	//addBus(Bus)
	canArrive(Bus) bool
	notifyAboutDeparture()
}
