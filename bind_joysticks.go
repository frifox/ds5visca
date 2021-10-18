package main

import "github.com/frifox/ds5"

func joysticks() {
	ds.Axis.Left.OnChange = func(j ds5.Joystick) {
		cam.Zoom.Z = j.Y
		cam.Apply(&cam.Zoom)
	}
	ds.Axis.Right.OnChange = func(j ds5.Joystick) {
		cam.Move.X = j.X
		cam.Move.Y = j.Y
		cam.Apply(&cam.Move)
	}
}
