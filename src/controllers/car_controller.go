package controllers

import (
	"juego/src/models"
	"juego/src/views"
	"math/rand"
	"time"
)

type CarController struct {
	Car        *models.Car
	Parking    *models.Parking
	CarView    *views.CarView
	CarManager *models.CarManager
	DoorChan   chan struct{}
	PathChan   chan struct{}
}

func NewCarController(car *models.Car, parking *models.Parking, carManager *models.CarManager, doorChan chan struct{}, pathChan chan struct{}) *CarController {
	return &CarController{
		Car:        car,
		Parking:    parking,
		CarManager: carManager,
		DoorChan:   doorChan,
		PathChan:   pathChan,
	}
}

// Start inicia el ciclo de vida del carro.
func (cicle *CarController) Start() {
	cicle.CarManager.AddCar(cicle.Car)

	cicle.Enqueue()

	spot := cicle.Parking.GetAvailableSpot()

	cicle.Park(spot)

	time.Sleep(time.Second * time.Duration(rand.Intn(15)+20))

	cicle.LeaveSpot()
	cicle.Parking.ReleaseSpot(spot)

	cicle.Leave(spot)

	// Los autos que salen adquieren el PathChan para tener prioridad
	<-cicle.PathChan
	cicle.ExitDoor()
	cicle.PathChan <- struct{}{}

	cicle.GoAway()
	cicle.CarManager.RemoveCar(cicle.Car)
}

// Implementación de los métodos para mover el carro.
func (MoveCar *CarController) Enqueue() {
	MoveCar.Parking.CarsList.Enqueue(MoveCar.Car)

	minY := 45.0
	spacing := 50.0

	for MoveCar.Car.Y > minY {
		carAhead := MoveCar.Parking.CarsList.GetCarAhead(MoveCar.Car)
		canMove := true

		if carAhead != nil {
			_, aheadY := carAhead.GetPosition()
			ccY := MoveCar.Car.Y

			if ccY-aheadY < spacing {
				canMove = false
			}
		}

		if canMove {
			MoveCar.Car.SetDirection(0, -1)
			MoveCar.Car.Move(0, -1)
		} else {
			MoveCar.Car.SetDirection(0, -1)
		}
		time.Sleep(10 * time.Millisecond)
	}

	// Acceso a la entrada
	<-MoveCar.DoorChan
	defer func() { MoveCar.DoorChan <- struct{}{} }()

	// Acceso al camino (compartido)
	<-MoveCar.PathChan
	defer func() { MoveCar.PathChan <- struct{}{} }()

	// Pasar por la entrada viendo colisiones
	MoveCar.JoinDoor()

	MoveCar.Parking.CarsList.RemoveCar(MoveCar.Car)
}

func (carDistance *CarController) JoinDoor() {
	minDistance := 50.0
	for carDistance.Car.X < 355 {
		canMove := true
		// Verificar si hay colisión
		for _, otherCar := range carDistance.CarManager.GetCars() {
			if otherCar != carDistance.Car {
				otherX, otherY := otherCar.GetPosition()
				if carDistance.Car.Y == otherY && carDistance.Car.X < otherX && otherX-carDistance.Car.X < minDistance {
					canMove = false
					break
				}
			}
		}
		if canMove {
			carDistance.Car.SetDirection(1, 0)
			carDistance.Car.Move(1, 0)
		} else {
			carDistance.Car.SetDirection(1, 0)
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func (carDirection *CarController) Park(spot *models.ParkingSpot) {
	for _, direction := range spot.GetDirectionsForParking() {
		carDirection.move(direction)
	}
}

func (carfront *CarController) LeaveSpot() {
	carfront.Car.SetDirection(0, -1)
	carfront.Car.Move(0, -30)
}

func (car *CarController) Leave(spot *models.ParkingSpot) {
	for _, direction := range spot.GetDirectionsForLeaving() {
		car.move(direction)
	}
}

func (carDistanceColi *CarController) ExitDoor() {
	minDistance := 50.0
	for carDistanceColi.Car.X > 300 {
		canMove := true
		for _, otherCar := range carDistanceColi.CarManager.GetCars() {
			if otherCar != carDistanceColi.Car {
				otherX, otherY := otherCar.GetPosition()
				if carDistanceColi.Car.Y == otherY && carDistanceColi.Car.X > otherX && carDistanceColi.Car.X-otherX < minDistance {
					canMove = false
					break
				}
			}
		}
		if canMove {
			carDistanceColi.Car.SetDirection(-1, 0)
			carDistanceColi.Car.Move(-1, 0)
		} else {
			carDistanceColi.Car.SetDirection(-1, 0)
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func (cc *CarController) GoAway() {
	cc.Car.SetDirection(-1, 0)
	for cc.Car.X > -20 {
		cc.Car.Move(-1, 0)
		time.Sleep(5 * time.Millisecond)
	}
	cc.CarManager.RemoveCar(cc.Car)
}

func (moveCarDirections *CarController) move(direction *models.ParkingSpotDirection) {
	minDistance := 50.0
	for {
		var canMove bool = true
		var dx, dy float64

		switch direction.Direction {
		case "left":
			if moveCarDirections.Car.X <= direction.Point {
				return
			}
			dx, dy = -1, 0
			for _, otherCar := range moveCarDirections.CarManager.GetCars() {
				if otherCar != moveCarDirections.Car {
					otherX, otherY := otherCar.GetPosition()
					if moveCarDirections.Car.Y == otherY && moveCarDirections.Car.X > otherX && moveCarDirections.Car.X-otherX < minDistance {
						canMove = false
						break
					}
				}
			}
		case "right":
			if moveCarDirections.Car.X >= direction.Point {
				return
			}
			dx, dy = 1, 0
			for _, otherCar := range moveCarDirections.CarManager.GetCars() {
				if otherCar != moveCarDirections.Car {
					otherX, otherY := otherCar.GetPosition()
					if moveCarDirections.Car.Y == otherY && moveCarDirections.Car.X < otherX && otherX-moveCarDirections.Car.X < minDistance {
						canMove = false
						break
					}
				}
			}
		case "up":
			if moveCarDirections.Car.Y <= direction.Point {
				return
			}
			dx, dy = 0, -1
			for _, otherCar := range moveCarDirections.CarManager.GetCars() {
				if otherCar != moveCarDirections.Car {
					otherX, otherY := otherCar.GetPosition()
					if moveCarDirections.Car.X == otherX && moveCarDirections.Car.Y > otherY && moveCarDirections.Car.Y-otherY < minDistance {
						canMove = false
						break
					}
				}
			}
		case "down":
			if moveCarDirections.Car.Y >= direction.Point {
				return
			}
			dx, dy = 0, 1
			for _, otherCar := range moveCarDirections.CarManager.GetCars() {
				if otherCar != moveCarDirections.Car {
					otherX, otherY := otherCar.GetPosition()
					if moveCarDirections.Car.X == otherX && moveCarDirections.Car.Y < otherY && otherY-moveCarDirections.Car.Y < minDistance {
						canMove = false
						break
					}
				}
			}
		}

		if canMove {
			moveCarDirections.Car.SetDirection(dx, dy)
			moveCarDirections.Car.Move(dx, dy)
		} else {
			// Mantener la dirección actual
			moveCarDirections.Car.SetDirection(dx, dy)
		}
		time.Sleep(10 * time.Millisecond)
	}
}
