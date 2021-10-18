# DualSense Controller + PTZOptics VISCA

Control a PTZOptics camera via VISCA over Serial or IP using a PlayStation 5 controller - DualSense.

This projects makes use of 2 packages I wrote to help me interface with DS5 and VISCA camera, [frifox/ds5](https://github.com/frifox/ds5) and [frifox/visca](https://github.com/frifox/visca).

DS5 controller will work over Bluetooth and USB.

For lowest latency, use controller over USB and camera over Serial. When testing inquiry command roundtrip times, VISCA over IP averaged 100ms and VISCA over Serial averaged 20ms.

## Controls

* Shape buttons recall presets 1-4. Long-pressing those shapes will save current position as that preset.
* Right joystic controls Pan/Tilt, capped at 50% max speed at the extremes
* Left joystic Y axis controls Zoom
* DPad keys trigger up/down/left/right motion for 20ms, for precise Pan/Tilt adjustments
* Touchpad control OSD. Press on left side for Menu/Back and press right side for Enter.
* When OSD is open, DPad keys act as navigation keys
* R1 triggers Focus
* Controller battery is reflected via LED on the controller.
  * Green = 50% or higher
  * Yellow = less than 50%
  * Red = less than 10%
* The 5 LEDs below touchpad identify battery levels 0% - 40%, each LED = 10% increment 