package main

import (
	"fmt"
	"github.com/frifox/ds5"
	"time"
)

var ds *ds5.Device

func runPS5() {
	ds.LightBar.Set(255, 255, 255)

	ds.Axis.Left.DeadZone = 10.0 / 128
	ds.Axis.Right.DeadZone = 10.0 / 128

	ds.Axis.Left.InvertY = true
	ds.Axis.Right.InvertY = true

	ds.Battery.OnChange = updateBattery
	ds.Buttons.Mute.OnKeyDown = updateBattery
	ds.Bus.OnChange = func(b string) {
		fmt.Printf("Bus is %s\n", b)
	}

	for {
		fmt.Printf("Looking for DS5\n")
		for {
			_ = ds.Find()
			if ds.Found() {
				fmt.Printf("DS5 found\n")
				break
			}
			time.Sleep(time.Second)
		}

		fmt.Printf("Running DS5\n")
		go ds.Run()

		<-ds.Done()
		fmt.Printf("DS5 died. Disconnected?\n")
	}
}

func updateBattery() {
	battery := &ds.Battery
	lights := &ds.LightBar
	leds := &ds.PlayerLEDs
	mic := &ds.Mic

	fmt.Printf("[Battery] %s (%d%%)\n", battery.Status, battery.Percent)

	// charge plug indicator
	switch battery.Status {
	case "Charging":
		mic.LED = true
	default:
		mic.LED = false
	}

	switch p := battery.Percent; {

	//case p >= 100:
	//	lights.SetGreen()
	//	leds.SetBar(5)
	//case p >= 90:
	//	lights.SetGreen()
	//	leds.SetBar(4)
	//case p >= 80:
	//	lights.SetGreen()
	//	leds.SetBar(3)
	//case p >= 70:
	//	lights.SetGreen()
	//	leds.SetBar(2)
	//case p >= 60:
	//	lights.SetGreen()
	//	leds.SetBar(1)
	case p >= 50:
		lights.SetGreen()
		leds.SetBar(0)

	case p >= 40:
		lights.SetYellow()
		leds.SetBar(4)
	case p >= 30:
		lights.SetYellow()
		leds.SetBar(3)
	case p >= 20:
		lights.SetYellow()
		leds.SetBar(2)

	case p >= 10:
		lights.SetOrange()
		leds.SetBar(1)

	case p > 0:
		lights.SetRed()
		leds.SetBar(0)
	}
	ds.ApplyProps()
}

func doHapticFeedback() {
	ds.Rumble.Right = 1
	ds.ApplyProps()

	time.Sleep(time.Millisecond * 100)

	ds.Rumble.Right = 0
	ds.ApplyProps()
}
