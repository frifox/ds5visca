package main

import "time"

func mapLeftButtons() {
	up := &ds.Buttons.DPadUp
	right := &ds.Buttons.DPadRight
	down := &ds.Buttons.DPadDown
	left := &ds.Buttons.DPadLeft

	duration := time.Millisecond * 20

	left.OnKeyDown = func() {
		cam.Move.X = -1.0 / 24
		cam.Apply(&cam.Move)

		time.Sleep(duration)

		cam.Move.X = 0
		cam.Apply(&cam.Move)
	}
	right.OnKeyDown = func() {
		cam.Move.X = 1.0 / 24
		cam.Apply(&cam.Move)

		time.Sleep(duration)

		cam.Move.X = 0
		cam.Apply(&cam.Move)
	}

	up.OnKeyDown = func() {
		cam.Move.Y = 1.0 / 24
		cam.Apply(&cam.Move)

		time.Sleep(duration)

		cam.Move.Y = 0
		cam.Apply(&cam.Move)
	}

	down.OnKeyDown = func() {
		cam.Move.Y = -1.0 / 24
		cam.Apply(&cam.Move)

		time.Sleep(duration)

		cam.Move.Y = 0
		cam.Apply(&cam.Move)
	}
}
