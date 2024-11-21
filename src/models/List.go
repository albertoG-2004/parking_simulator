package models

import (
	"sync"
)

// CarQueue representa una cola de autos.
type CarInListCars struct {
	mu   sync.Mutex
	cars []*Car // Slice de autos en la cola.
}

// NewCarQueue crea una nueva instancia de CarQueue.
func NewCarQueue() *CarInListCars {
	return &CarInListCars{
		cars: []*Car{},
	}
}

// Enqueue añade un auto a la cola.
func (carList *CarInListCars) Enqueue(car *Car) {
	carList.mu.Lock()
	defer carList.mu.Unlock()
	carList.cars = append(carList.cars, car)
}

// Dequeue elimina y devuelve el primer auto de la cola.
func (carList *CarInListCars) Dequeue() *Car {
	carList.mu.Lock()
	defer carList.mu.Unlock()
	if len(carList.cars) == 0 {
		return nil
	}
	car := carList.cars[0]
	carList.cars = carList.cars[1:]
	return car
}

// GetPositionInQueue devuelve la posición de un auto en la cola.
func (carList *CarInListCars) GetPositionInQueue(car *Car) int {
	carList.mu.Lock()
	defer carList.mu.Unlock()
	for i, c := range carList.cars {
		if c == car {
			return i
		}
	}
	return -1
}

// GetCarAhead devuelve el auto que está delante en la cola.
func (carList *CarInListCars) GetCarAhead(car *Car) *Car {
	position := carList.GetPositionInQueue(car)
	if position > 0 {
		return carList.cars[position-1]
	}
	return nil
}

// RemoveCar elimina un auto de la cola.
func (carList *CarInListCars) RemoveCar(car *Car) {
	carList.mu.Lock()
	defer carList.mu.Unlock()
	for i, c := range carList.cars {
		if c == car {
			carList.cars = append(carList.cars[:i], carList.cars[i+1:]...)
			break
		}
	}
}
