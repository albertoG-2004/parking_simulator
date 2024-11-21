package models

import "sync"

// CarManager gestiona una colección de autos.
type CarManager struct {
	mu   sync.Mutex
	Cars []*Car
}

// NewCarManager crea una nueva instancia de CarManager.
func NewCarManager() *CarManager {
	return &CarManager{
		Cars: []*Car{},
	}
}

// AddCar añade un auto al gestor.
func (carAddGes *CarManager) AddCar(car *Car) {
	carAddGes.mu.Lock()
	defer carAddGes.mu.Unlock()
	carAddGes.Cars = append(carAddGes.Cars, car)
}

// RemoveCar elimina un auto del gestor.
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

// GetCars devuelve la lista de autos.
func (Listdevolution *CarManager) GetCars() []*Car {
	Listdevolution.mu.Lock()
	defer Listdevolution.mu.Unlock()
	carsCopy := make([]*Car, len(Listdevolution.Cars))
	copy(carsCopy, Listdevolution.Cars)
	return carsCopy
}
