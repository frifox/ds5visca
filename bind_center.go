package main

import "log"

func mapCenterButtons() {
	ds.Buttons.PS.OnKeyDown = func() {
		err := ds.BTDisconnect()
		if err != nil {
			log.Printf("ERR BTDisconnect: %v\n", err)
		}
	}
}
