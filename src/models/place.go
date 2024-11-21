package models

// Parking representa el estacionamiento.
type Parking struct {
	Spots          []*ParkingSpot
	CarsList      *CarInListCars
	AvailableSpots chan *ParkingSpot
}

// NewParking crea una nueva instancia de Parking.
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

// GetAvailableSpot obtiene un lugar disponible.
func (placeNull *Parking) GetAvailableSpot() *ParkingSpot {
	return <-placeNull.AvailableSpots
}

// ReleaseSpot libera un lugar ocupado.
func (placeNull *Parking) ReleaseSpot(spot *ParkingSpot) {
	placeNull.AvailableSpots <- spot
}
