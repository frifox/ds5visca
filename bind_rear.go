package main

import "fmt"

func mapRearButtons() {
	ds.Buttons.R1.OnKeyDown = func() {
		fmt.Println("[Focus]")
		cam.Apply(&cam.Focus)
	}
}
