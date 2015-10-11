package main

import (
	"math"
	"time"

	"github.com/Ganners/gorender/window"
)

// Small test program to get things running
func main() {
	window := window.NewWindow("My Application", 800, 600)

	// Create window to draw on
	window.Create()
	defer window.Destroy()

	// Calculate midpoints
	midX := math.Floor(float64(window.Width) / 2.0)
	midY := math.Floor(float64(window.Height) / 2.0)

	maxRadius := math.Min(float64(window.Height), float64(window.Width))

	for frame := 0.0; ; frame++ {

		window.Backbuffer.Fill(0xFF000000)
		radius := float64(int(frame)%int(maxRadius)) / 2.0

		// Draw something weird on-screen
		for i := 0.0; i < 5.0; i++ {
			for theta := 0.0; theta < 2.0*math.Pi; theta += 0.005 {

				y := int(math.Floor(midY - (radius-(i+theta*20.0))*math.Sin(theta*i+(frame))))
				x := int(math.Floor(midX - (radius-(i+theta*50.0))*math.Cos(theta*i+(frame))))

				window.Backbuffer.DrawPixel(
					x, y, uint32(0xFF0000FF)|((uint32(frame)%255)<<8))
			}
			window.Update()
		}
		time.Sleep(time.Millisecond * (1000 / 30))
	}
}
