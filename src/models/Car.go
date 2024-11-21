package models

import (
	"math"
	"math/rand"
	"sync"
	"time"
)

type Car struct {
	mu        sync.Mutex
	X         float64
	Y         float64
	DX        float64
	DY        float64
	ModelPath string
	observers []Observer
}

func NewCar() *Car {
	modelPaths := []string{
		"src/assets/brownCar.png",
		"src/assets/green.png",
		"src/assets/orangeCar.png",
	}
	rand.Seed(time.Now().UnixNano())
	modelPath := modelPaths[rand.Intn(len(modelPaths))]

	return &Car{
		X:         200,
		Y:         600,
		DX:        0,
		DY:        -1,
		ModelPath: modelPath,
		observers: []Observer{},
	}
}

func (carObserAdd *Car) RegisterObserver(o Observer) {
	carObserAdd.mu.Lock()
	defer carObserAdd.mu.Unlock()
	carObserAdd.observers = append(carObserAdd.observers, o)
}

func (CarObserDelete *Car) RemoveObserver(o Observer) {
	CarObserDelete.mu.Lock()
	defer CarObserDelete.mu.Unlock()
	for i, observer := range CarObserDelete.observers {
		if observer == o {
			CarObserDelete.observers = append(CarObserDelete.observers[:i], CarObserDelete.observers[i+1:]...)
			break
		}
	}
}

// NotifyObservers notifica a todos los observadores sobre un cambio.
func (carNotifyObserver *Car) NotifyObservers() {
	carNotifyObserver.mu.Lock()
	observers := make([]Observer, len(carNotifyObserver.observers))
	copy(observers, carNotifyObserver.observers)
	carNotifyObserver.mu.Unlock()

	for _, observer := range observers {
		observer.Update(carNotifyObserver)
	}
}

func (carUpdatePosition *Car) Move(dx, dy float64) {
	carUpdatePosition.mu.Lock()
	carUpdatePosition.X += dx
	carUpdatePosition.Y += dy
	carUpdatePosition.mu.Unlock()
	carUpdatePosition.NotifyObservers()
}

func (carResDirection *Car) SetDirection(dx, dy float64) {
	carResDirection.mu.Lock()
	carResDirection.DX = dx
	carResDirection.DY = dy
	carResDirection.mu.Unlock()
	carResDirection.NotifyObservers()
}

func (carGetPositionActually *Car) GetPosition() (float64, float64) {
	carGetPositionActually.mu.Lock()
	defer carGetPositionActually.mu.Unlock()
	return carGetPositionActually.X, carGetPositionActually.Y
}

func (carUltimateMove *Car) GetDirection() (float64, float64) {
	carUltimateMove.mu.Lock()
	defer carUltimateMove.mu.Unlock()
	return carUltimateMove.DX, carUltimateMove.DY
}

func (carDirectionForm *Car) GetDirectionName() string {
	carDirectionForm.mu.Lock()
	dx := carDirectionForm.DX
	dy := carDirectionForm.DY
	carDirectionForm.mu.Unlock()

	if dx == 0 && dy == 0 {
		return "up"
	}

	angle := math.Atan2(dy, dx) * (180 / math.Pi)

	if angle >= -45 && angle <= 45 {
		return "right"
	} else if angle > 45 && angle < 135 {
		return "up"
	} else if angle >= 135 || angle <= -135 {
		return "left"
	} else {
		return "down"
	}
}

func (carPositionRessetY *Car) SetX(x float64) {
	carPositionRessetY.mu.Lock()
	defer carPositionRessetY.mu.Unlock()
	carPositionRessetY.X = x
	carPositionRessetY.NotifyObservers()
}

func (carPositionsetX *Car) SetY(y float64) {
	carPositionsetX.mu.Lock()
	defer carPositionsetX.mu.Unlock()
	carPositionsetX.Y = y
	carPositionsetX.NotifyObservers()
}
