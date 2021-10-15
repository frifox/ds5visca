package main

func joysticks() {
	ds.Axis.Left.OnChange = func(x, y float64) {
		cam.Zoom.Z = y
		cam.Apply(&cam.Zoom)
	}
	ds.Axis.Right.OnChange = func(x, y float64) {
		cam.Move.X = x
		cam.Move.Y = y
		cam.Apply(&cam.Move)
	}
}
