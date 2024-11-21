package models

import "sync"

type CarManager struct {
	mu   sync.Mutex
	Cars []*Car
}

func NewCarManager() *CarManager {
	return &CarManager{
		Cars: []*Car{},
	}
}

// AddCar añade un auto al gestor (carManager).
func (carAddGes *CarManager) AddCar(car *Car) {
	carAddGes.mu.Lock()
	defer carAddGes.mu.Unlock()
	carAddGes.Cars = append(carAddGes.Cars, car)
}

// RemoveCar elimina un auto del gestor (carManager).
func (carDeleteGes *CarManager) RemoveCar(car *Car) {
	carDeleteGes.mu.Lock()
	defer carDeleteGes.mu.Unlock()
	for i, c := range carDeleteGes.Cars {
		if c == car {
			carDeleteGes.Cars = append(carDeleteGes.Cars[:i], carDeleteGes.Cars[i+1:]...)
			break
		}
	}
}

func (Listdevolution *CarManager) GetCars() []*Car {
	Listdevolution.mu.Lock()
	defer Listdevolution.mu.Unlock()
	carsCopy := make([]*Car, len(Listdevolution.Cars))
	copy(carsCopy, Listdevolution.Cars)
	return carsCopy
}
