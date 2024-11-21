package views

import (
	"juego/src/models"

	"github.com/oakmound/oak/v4/render"
	"github.com/oakmound/oak/v4/scene"
)

// ParkingView representa la vista del estacionamiento.
type ParkingView struct {
	Parking *models.Parking
	Context *scene.Context
}

// NewParkingView crea una nueva instancia de ParkingView.
func NewParkingView(parking *models.Parking, ctx *scene.Context) *ParkingView {
	pv := &ParkingView{
		Parking: parking,
		Context: ctx,
	}
	pv.setupScene()
	return pv
}

// setupScene configura visualmente el estacionamiento.
func (pv *ParkingView) setupScene() {
	// Cargar la imagen de fondo como sprite.
	backgroundImage, err := render.LoadSprite("src/assets/parkingNull.jpg")
	if err != nil {
		panic("Error al cargar la imagen de fondo del estacionamiento: " + err.Error())
	}
	// Establecer posición y dibujar la imagen de fondo en el contexto.
	backgroundImage.SetPos(0, 0)    // Establecer posición inicial.
	render.Draw(backgroundImage, 0) // Dibujar en la capa más baja.

}
