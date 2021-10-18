package main

import (
	"github.com/JamesLMilner/pip-go"
	"github.com/frifox/ds5"
	"github.com/frifox/visca"
)

func mapTouchpad() {
	// Touchpad button will be different based on where you click
	points := map[string]pip.Point{
		"TL": {0, 0},
		"TR": {1920, 0},
		"BR": {1920, 1080},
		"BL": {0, 1080},

		"T": {1920 / 2, 0},
		"M": {1920 / 2, 1080 / 2},
		"B": {1920 / 2, 1080},
	}

	left := pip.Polygon{
		Points: []pip.Point{points["TL"], points["T"], points["B"], points["BL"], points["TL"]},
	}
	right := pip.Polygon{
		Points: []pip.Point{points["T"], points["TR"], points["BR"], points["B"], points["T"]},
	}

	var button visca.Command

	lastID := uint8(0)
	ds.Touchpad.Touch1.OnActive = func(t ds5.Touch) {
		if t.ID == lastID {
			return
		}

		// where are we touching?
		point := pip.Point{X: float64(t.X), Y: float64(t.Y)}

		// left half = OSDToggle
		if pip.PointInPolygon(point, left) && button != &cam.OSDToggle {
			button = &cam.OSDToggle
		}

		// right half = OSDEnter
		if pip.PointInPolygon(point, right) && button != &cam.OSDEnter {
			button = &cam.OSDEnter
		}
	}

	// trigger selected button on touchpad press
	ds.Buttons.Touchpad.OnKeyDown = func() {
		cam.Apply(button)
	}

}
