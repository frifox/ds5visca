package main

func rear() {
	ds.Buttons.R1.OnKeyDown = func() {
		cam.Apply(&cam.Focus)
	}
}
