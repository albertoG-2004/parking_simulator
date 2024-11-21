package models

import (
	"github.com/oakmound/oak/v4/alg/floatgeom"
)

// Constantes para direcciones
const (
	DirectionLeft  = "left"
	DirectionRight = "right"
	DirectionUp    = "up"
	DirectionDown  = "down"
)

// ParkingSpotDirection representa una dirección específica y un punto asociado para maniobrar en un estacionamiento.
type ParkingSpotDirection struct {
	Direction string  // Dirección del movimiento.
	Point     float64 // Punto asociado con la dirección.
}

// NewParkingSpotDirection crea y devuelve una nueva instancia de ParkingSpotDirection.
func NewParkingSpotDirection(direction string, point float64) *ParkingSpotDirection {
	return &ParkingSpotDirection{
		Direction: direction,
		Point:     point,
	}
}

// ParkingSpot representa un lugar de estacionamiento.
type ParkingSpot struct {
	Area                 *floatgeom.Rect2        // Área que delimita el lugar de estacionamiento.
	DirectionsForParking []*ParkingSpotDirection // Direcciones para estacionar en este lugar.
	DirectionsForLeaving []*ParkingSpotDirection // Direcciones para salir de este lugar.
	Number               int                     // Número del lugar de estacionamiento.
	IsAvailable          bool                    // Estado de disponibilidad del lugar.
}

// NewParkingSpot crea y devuelve un nuevo lugar de estacionamiento con los parámetros especificados.
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

// GetDirectionsForParking genera y devuelve las direcciones necesarias para estacionar en función de la fila.
func GetDirectionsForParking(row int, x, y float64) []*ParkingSpotDirection {
	var directions []*ParkingSpotDirection

	// Define la dirección principal basada en la fila.
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

	// Añade movimientos adicionales hacia la derecha y abajo.
	directions = append(directions,
		NewParkingSpotDirection(DirectionRight, x+5),
		NewParkingSpotDirection(DirectionDown, y+5),
	)

	return directions
}

// GetDirectionsForLeaving genera y devuelve las direcciones necesarias para salir del lugar de estacionamiento.
func GetDirectionsForLeaving() []*ParkingSpotDirection {
	// Añade las direcciones para salir del estacionamiento.
	return []*ParkingSpotDirection{
		NewParkingSpotDirection(DirectionRight, 600),
		NewParkingSpotDirection(DirectionUp, 15),
		NewParkingSpotDirection(DirectionLeft, 355),
	}
}

// Métodos de acceso y modificación:

// GetArea devuelve el área del lugar de estacionamiento.
func (parkDevolution *ParkingSpot) GetArea() *floatgeom.Rect2 {
	return parkDevolution.Area
}

// GetNumber devuelve el número de identificación del lugar de estacionamiento.
func (parkIdentification *ParkingSpot) GetNumber() int {
	return parkIdentification.Number
}

// GetDirectionsForParking devuelve las direcciones necesarias para estacionar en el lugar.
func (parkDirectionsNecesary *ParkingSpot) GetDirectionsForParking() []*ParkingSpotDirection {
	return parkDirectionsNecesary.DirectionsForParking
}

// GetDirectionsForLeaving devuelve las direcciones necesarias para salir del lugar.
func (parkDevolutionOut *ParkingSpot) GetDirectionsForLeaving() []*ParkingSpotDirection {
	return parkDevolutionOut.DirectionsForLeaving
}

// GetIsAvailable verifica si el lugar de estacionamiento está disponible.
func (carVerifyPlace *ParkingSpot) GetIsAvailable() bool {
	return carVerifyPlace.IsAvailable
}

// SetIsAvailable establece el estado de disponibilidad del lugar de estacionamiento.
func (carEstate *ParkingSpot) SetIsAvailable(isAvailable bool) {
	carEstate.IsAvailable = isAvailable
}

// Métodos adicionales para obtener las coordenadas del área:

// GetX devuelve la coordenada mínima X del lugar de estacionamiento.
func (carNotifyAll *ParkingSpot) GetX() float64 {
	return carNotifyAll.Area.Min.X()
}

// GetY devuelve la coordenada mínima Y del lugar de estacionamiento.
func (carCoorde *ParkingSpot) GetY() float64 {
	return carCoorde.Area.Min.Y()
}

// GetX2 devuelve la coordenada máxima X del lugar de estacionamiento.
func (carCoorX *ParkingSpot) GetX2() float64 {
	return carCoorX.Area.Max.X()
}

// GetY2 devuelve la coordenada máxima Y del lugar de estacionamiento.
func (carCoorY *ParkingSpot) GetY2() float64 {
	return carCoorY.Area.Max.Y()
}
