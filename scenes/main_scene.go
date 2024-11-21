package scenes

import (
	"juego/src/controllers"
	"juego/src/models"
	"juego/src/views"

	"github.com/oakmound/oak/v4/scene"
)

func NewParkingScene() *scene.Scene {
	return &scene.Scene{
		Start: func(ctx *scene.Context) {
			spots := []*models.ParkingSpot{
				models.NewParkingSpot(380, 70, 460, 130, 1, 1),
				models.NewParkingSpot(485, 70, 565, 130, 1, 2),
				models.NewParkingSpot(590, 70, 670, 130, 1, 3),
				models.NewParkingSpot(695, 70, 775, 130, 1, 4),
				models.NewParkingSpot(800, 70, 880, 130, 1, 5),

				models.NewParkingSpot(380, 160, 460, 220, 2, 6),
				models.NewParkingSpot(485, 160, 565, 220, 2, 7),
				models.NewParkingSpot(590, 160, 670, 220, 2, 8),
				models.NewParkingSpot(695, 160, 775, 220, 2, 9),
				models.NewParkingSpot(800, 160, 880, 220, 2, 10),

				models.NewParkingSpot(380, 450, 460, 510, 3, 11),
				models.NewParkingSpot(485, 450, 565, 510, 3, 12),
				models.NewParkingSpot(590, 450, 670, 510, 3, 13),
				models.NewParkingSpot(695, 450, 775, 510, 3, 14),
				models.NewParkingSpot(800, 450, 880, 510, 3, 15),

				models.NewParkingSpot(380, 540, 460, 600, 4, 16),
				models.NewParkingSpot(485, 540, 565, 600, 4, 17),
				models.NewParkingSpot(590, 540, 670, 600, 4, 18),
				models.NewParkingSpot(695, 540, 775, 600, 4, 19),
				models.NewParkingSpot(800, 540, 880, 600, 4, 20),

			}

			parking := models.NewParking(spots)
			parkingController := controllers.NewParkingController(parking)
			parkingView := views.NewParkingView(parking, ctx)
			parkingController.View = parkingView

			parkingController.StartCarGeneration(ctx)
		},
	}
}
