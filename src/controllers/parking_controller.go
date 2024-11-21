package controllers

import (
	"juego/src/models"
	"juego/src/views"
	"math/rand"
	"time"

	"github.com/oakmound/oak/v4/scene"
)

// parking_controller maneja la lógica del estacionamiento.
type ParkingController struct {
	Parking    *models.Parking
	View       *views.ParkingView
	DoorChan   chan struct{}
	PathChan   chan struct{}
	CarManager *models.CarManager
}

// NewParkingController crea instancia de ParkingController.
func NewParkingController(parking *models.Parking) *ParkingController {
	doorChan := make(chan struct{}, 1)
	doorChan <- struct{}{}

	pathChan := make(chan struct{}, 1)
	pathChan <- struct{}{}

	return &ParkingController{
		Parking:    parking,
		DoorChan:   doorChan,
		PathChan:   pathChan,
		CarManager: models.NewCarManager(),
	}
}

// StartCarGeneration inicia el proceso de generación de autos.
func (pc *ParkingController) StartCarGeneration(ctx *scene.Context) {
	const maxCars = 100
	go func() {
		for i := 0; i < maxCars; i++ {
			car := models.NewCar()
			carController := NewCarController(car, pc.Parking, pc.CarManager, pc.DoorChan, pc.PathChan)
			carView := views.NewCarView(car, ctx)
			carController.CarView = carView
			go carController.Start()
			time.Sleep(time.Second * time.Duration(rand.Intn(2)+1))
		}
	}()
}
