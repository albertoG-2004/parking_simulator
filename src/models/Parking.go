package models

import (
	"github.com/oakmound/oak/v4/alg/floatgeom"
)

const (
	DirectionLeft  = "left"
	DirectionRight = "right"
	DirectionUp    = "up"
	DirectionDown  = "down"
)

type ParkingSpotDirection struct {
	Direction string  // Dirección del movimiento.
	Point     float64 // Punto asociado con la dirección.
}

func NewParkingSpotDirection(direction string, point float64) *ParkingSpotDirection {
	return &ParkingSpotDirection{
		Direction: direction,
		Point:     point,
	}
}

// Lugar de estacionamiento
type ParkingSpot struct {
	Area                 *floatgeom.Rect2
	DirectionsForParking []*ParkingSpotDirection
	DirectionsForLeaving []*ParkingSpotDirection
	Number               int
	IsAvailable          bool
}

func NewParkingSpot(x, y, x2, y2 float64, row, number int) *ParkingSpot {
	area := floatgeom.NewRect2(x, y, x2, y2)
	directionsForParking := GetDirectionsForParking(row, x, y)
	directionsForLeaving := GetDirectionsForLeaving()

	return &ParkingSpot{
		Area:                 &area,
		DirectionsForParking: directionsForParking,
		DirectionsForLeaving: directionsForLeaving,
		Number:               number,
		IsAvailable:          true,
	}
}

func GetDirectionsForParking(row int, x, y float64) []*ParkingSpotDirection {
	var directions []*ParkingSpotDirection

	switch row {
	case 1:
		directions = append(directions, NewParkingSpotDirection(DirectionDown, 45))
	case 2:
		directions = append(directions, NewParkingSpotDirection(DirectionDown, 135))
	case 3:
		directions = append(directions, NewParkingSpotDirection(DirectionDown, 225))
	case 4:
		directions = append(directions, NewParkingSpotDirection(DirectionDown, 315))
	}

	directions = append(directions,
		NewParkingSpotDirection(DirectionRight, x+5),
		NewParkingSpotDirection(DirectionDown, y+5),
	)

	return directions
}

func GetDirectionsForLeaving() []*ParkingSpotDirection {
	// Añade las direcciones para salir del estacionamiento.
	return []*ParkingSpotDirection{
		NewParkingSpotDirection(DirectionRight, 600),
		NewParkingSpotDirection(DirectionUp, 15),
		NewParkingSpotDirection(DirectionLeft, 355),
	}
}

func (parkDevolution *ParkingSpot) GetArea() *floatgeom.Rect2 {
	return parkDevolution.Area
}

func (parkIdentification *ParkingSpot) GetNumber() int {
	return parkIdentification.Number
}

func (parkDirectionsNecesary *ParkingSpot) GetDirectionsForParking() []*ParkingSpotDirection {
	return parkDirectionsNecesary.DirectionsForParking
}

func (parkDevolutionOut *ParkingSpot) GetDirectionsForLeaving() []*ParkingSpotDirection {
	return parkDevolutionOut.DirectionsForLeaving
}

func (carVerifyPlace *ParkingSpot) GetIsAvailable() bool {
	return carVerifyPlace.IsAvailable
}

func (carEstate *ParkingSpot) SetIsAvailable(isAvailable bool) {
	carEstate.IsAvailable = isAvailable
}

func (carNotifyAll *ParkingSpot) GetX() float64 {
	return carNotifyAll.Area.Min.X()
}

func (carCoorde *ParkingSpot) GetY() float64 {
	return carCoorde.Area.Min.Y()
}

func (carCoorX *ParkingSpot) GetX2() float64 {
	return carCoorX.Area.Max.X()
}

func (carCoorY *ParkingSpot) GetY2() float64 {
	return carCoorY.Area.Max.Y()
}
