package models

type Parking struct {
	Spots          []*ParkingSpot
	CarsList      *CarInListCars
	AvailableSpots chan *ParkingSpot
}

func NewParking(spots []*ParkingSpot) *Parking {
	availableSpots := make(chan *ParkingSpot, len(spots))
	for _, spot := range spots {
		availableSpots <- spot
	}

	return &Parking{
		Spots:          spots,
		CarsList:      NewCarQueue(),
		AvailableSpots: availableSpots,
	}
}

func (placeNull *Parking) GetAvailableSpot() *ParkingSpot {
	return <-placeNull.AvailableSpots
}

func (placeNull *Parking) ReleaseSpot(spot *ParkingSpot) {
	placeNull.AvailableSpots <- spot
}