package main

import (
	"fmt"
	"time"
)
import "github.com/frifox/visca"

var cam *visca.Device

func runCamera() {
	for {
		fmt.Printf("Looking for camera at %s\n", cam.Path)
		for {
			_ = cam.Find()
			if cam.Found() {
				fmt.Printf("Camera found\n")
				break
			}
			time.Sleep(time.Second)
		}

		fmt.Printf("Running camera\n")
		go cam.Run()

		<-cam.Done()
		fmt.Printf("Camera quit\n")
	}
}
