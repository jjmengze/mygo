package main

import "github.com/jjmengze/mygo/pkg/design-pattern/mediator"

func main() {
	stationManager := mediator.NewStationManger()
	intercityBus := mediator.IntercityBus{
		Mediator: stationManager,
	}
	shuttleBus := mediator.ShuttleBus{
		Mediator: stationManager,
	}
	intercityBus.Arrive()
	shuttleBus.Arrive()
	intercityBus.Depart()
}
