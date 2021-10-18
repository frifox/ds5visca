package main

import "fmt"

func rear() {
	ds.Buttons.R1.OnKeyDown = func() {
		fmt.Println("[Focus]")
		cam.Apply(&cam.Focus)
	}
}
