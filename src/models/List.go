package models

import (
	"sync"
)

type CarInListCars struct {
	mu   sync.Mutex
	cars []*Car
}

func NewCarQueue() *CarInListCars {
	return &CarInListCars{
		cars: []*Car{},
	}
}

func (carList *CarInListCars) Enqueue(car *Car) {
	carList.mu.Lock()
	defer carList.mu.Unlock()
	carList.cars = append(carList.cars, car)
}

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

func (carList *CarInListCars) GetCarAhead(car *Car) *Car {
	position := carList.GetPositionInQueue(car)
	if position > 0 {
		return carList.cars[position-1]
	}
	return nil
}

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
