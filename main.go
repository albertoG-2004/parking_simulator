package main

import (
	"juego/scenes"

	"github.com/oakmound/oak/v4"
)

func main() {
	scene := scenes.NewParkingScene()
	oak.AddScene("parkingScene", *scene)
	err := oak.Init("parkingScene", func(c oak.Config) (oak.Config, error) {
		c.Screen.Width = 1300
		c.Screen.Height = 650
		c.Assets.ImagePath = "assets"
		return c, nil
	})
	if err != nil {
		panic(err)
	}
}
