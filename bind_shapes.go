package main

import (
	"context"
	"github.com/frifox/ds5"
)

func rightButtons() {
	rightButtonsCallSavePreset(1, &ds.Buttons.Square)
	rightButtonsCallSavePreset(2, &ds.Buttons.Cross)
	rightButtonsCallSavePreset(3, &ds.Buttons.Triangle)
	rightButtonsCallSavePreset(4, &ds.Buttons.Circle)
}

func rightButtonsCallSavePreset(id uint8, button *ds5.Button) {
	var ctx context.Context
	var cancel context.CancelFunc

	call := &cam.CallPreset
	save := &cam.SavePreset

	button.OnKeyDown = func() {
		ctx, cancel = context.WithCancel(context.Background())
	}
	button.OnKeyUp = func() {
		if ctx.Err() != nil {
			return
		}
		cancel()

		call.ID = id
		cam.Apply(call)
	}
	button.OnLongPress = func() {
		if ctx.Err() != nil {
			return // c
		}
		cancel()
		go doHapticFeedback()

		save.ID = id
		cam.Apply(save)
	}

}
