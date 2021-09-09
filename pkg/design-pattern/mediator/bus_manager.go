package mediator

type stationManager struct {
	isPlatformFree bool
	buses          []Bus
}

func NewStationManger() *stationManager {
	return &stationManager{
		isPlatformFree: true,
	}
}

//func (s *stationManager) addBus(b Bus) bool {
//	b.set
//	s.buses = append(s.buses, b)
//}

func (s *stationManager) canArrive(b Bus) bool {
	if s.isPlatformFree {
		s.isPlatformFree = false
		return true
	}
	s.buses = append(s.buses, b)
	return false
}

func (s *stationManager) notifyAboutDeparture() {
	if !s.isPlatformFree {
		s.isPlatformFree = true
	}
	if len(s.buses) > 0 {
		permittedArrivalBus := s.buses[0]
		s.buses = s.buses[1:]
		permittedArrivalBus.PermitArrival()
	}
}
