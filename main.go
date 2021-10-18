package main

import (
	"flag"
	"github.com/frifox/ds5"
	"github.com/frifox/visca"
)

var config = Config{}

type Config struct {
	CamPath string
}

func main() {
	flag.StringVar(&config.CamPath, "cam", "/dev/ttyAMA1", "tcp://ip:port || udp://ip:port || /dev/ttyAMAx")
	flag.Parse()

	ds = &ds5.Device{}
	cam = &visca.Device{
		Path: config.CamPath,
	}

	go runPS5()
	go runCamera()

	mapLeftButtons()
	mapCenterButtons()
	mapRightButtons()
	mapRearButtons()

	mapJoysticks()
	mapTouchpad()

	select {}
}
