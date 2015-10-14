package main

import (
	"math"
	"math/rand"

	"github.com/Ganners/gorender/backbuffer"
	"github.com/Ganners/gorender/window"
)

func main() {

	window := window.NewWindow("My Application", 800, 600)

	// Create window to draw on
	window.Create()
	defer window.Destroy()

	// Start starfield concurrently
	go func() {
		starField := NewStarField(
			2048, window.Width, window.Height,
			60.0,
			window.Backbuffer,
			func() { window.Update() })

		for {
			starField.Update(0.03)
		}
	}()

	// Launch application for 5 seconds
	haltChan := make(chan struct{}, 1)
	<-haltChan
}

// Vector stores our X, Y and Z coordinate values
type Vector3D struct {
	X float64
	Y float64
	Z float64
}

// Starfield is the structure for the program which handles calculation and
// drawing, it knows how to use a backbuffer and is told how to redraw
type StarField struct {
	stars      []Vector3D
	width      int
	height     int
	halfWidth  int
	halfHeight int
	backbuffer *backbuffer.Backbuffer
	redraw     func()
	time       int64
	fov        float64
}

// Constructs a new star field. Field of view should be in degrees
func NewStarField(
	numStars, displayWidth, displayHeight int,
	fieldOfView float64,
	backbuffer *backbuffer.Backbuffer,
	redraw func()) *StarField {

	starField := &StarField{
		stars:      make([]Vector3D, numStars),
		width:      displayWidth,
		height:     displayHeight,
		halfWidth:  displayWidth / 2,
		halfHeight: displayHeight / 2,
		backbuffer: backbuffer,
		redraw:     redraw,
		fov:        math.Tan((fieldOfView * 0.0174533) / 2),
	}

	// Position all of the stars
	starField.positionAllStars()

	return starField
}

func (sf *StarField) positionAllStars() {

	for i := 0; i < len(sf.stars); i++ {
		sf.positionStar(i)
	}
}

// Randomly positions a star in space
func (sf *StarField) positionStar(index int) {

	randX := (2.0 * (rand.Float64() - 0.5)) * 40.0
	randY := (2.0 * (rand.Float64() - 0.5)) * 40.0
	randZ := (rand.Float64() + 0.00001) * 40.0

	sf.stars[index].X = randX * float64(sf.halfWidth)
	sf.stars[index].Y = randY * float64(sf.halfHeight)
	sf.stars[index].Z = randZ
}

// Updates a frame of the starfield, handles drawing to backbuffer and
// redrawing window
func (sf *StarField) Update(delta float64) {

	if (sf.time % 10) == 0 {
		sf.backbuffer.Fill(0xFF000000)
	}

	for i := 0; i < len(sf.stars); i++ {

		// Move Z down to zero
		sf.stars[i].Z -= delta
		if sf.stars[i].Z <= 0 {
			sf.positionStar(i)
		}

		// Calculate X and Y pixel coordinates for perspective location
		x := int(sf.stars[i].X/(sf.stars[i].Z*sf.fov)) + sf.halfWidth
		y := int(sf.stars[i].Y/(sf.stars[i].Z*sf.fov)) + sf.halfHeight

		if (x < 0 || x >= sf.width-1) || (y < 0 || y >= sf.height-1) {
			// Restart if it goes off screen
			sf.positionStar(i)
		}

		// Draw a star with colour determined from Z position
		shade := (uint8(-sf.stars[i].Z) * 6)
		color := uint32(0xFF000000) | uint32(shade)<<16 |
			uint32(shade)<<8 | uint32(shade)

		sf.backbuffer.DrawPixel(x, y, color)
	}

	// Update canvas and time
	sf.time++
	sf.redraw()
}
