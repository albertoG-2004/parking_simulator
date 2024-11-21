package views

import (
	"juego/src/models"

	"github.com/oakmound/oak/v4/render"
	"github.com/oakmound/oak/v4/scene"
)

type ParkingView struct {
	Parking *models.Parking
	Context *scene.Context
}

func NewParkingView(parking *models.Parking, ctx *scene.Context) *ParkingView {
	pv := &ParkingView{
		Parking: parking,
		Context: ctx,
	}
	pv.setupScene()
	return pv
}

func (pv *ParkingView) setupScene() {
	backgroundImage, err := render.LoadSprite("src/assets/parkingNull.jpg")
	if err != nil {
		panic("Error al cargar la imagen de fondo del estacionamiento: " + err.Error())
	}

	backgroundImage.SetPos(0, 0)
	render.Draw(backgroundImage, 0)
}
