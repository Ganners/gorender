package backbuffer

import "errors"

// Stores all the necessary information for drawing pixels
type Backbuffer struct {
	width         int
	height        int
	bytesPerPixel int
	pixels        []byte
}

// Sets up a new backbuffer
func New(width, height, bytesPerPixel int, pixels []byte) *Backbuffer {
	return &Backbuffer{
		width:         width,
		height:        height,
		bytesPerPixel: bytesPerPixel,
		pixels:        pixels,
	}
}

// Draws a pixel at given coordinates, color is 0xAARRGGBB
func (bb *Backbuffer) DrawPixel(x, y int, color uint32) error {

	pos := ((y * bb.width) + x) * bb.bytesPerPixel

	// Stop unsafe operations
	if pos < 0 {
		return errors.New("Position cannot be negative")
	}
	if pos+3 >= len(bb.pixels)-1 {
		return errors.New("Pixel exceeds backbuffer, cannot set")
	}

	// Blue, Green, Red, Alpha
	bb.pixels[pos] = byte(uint8(color))
	bb.pixels[pos+1] = byte(uint8(color >> 8))
	bb.pixels[pos+2] = byte(uint8(color >> 16))
	bb.pixels[pos+3] = byte(uint8(color >> 24))

	return nil
}

// Fills the backbuffer with a particular color. Color is 0xAARRGGBB
func (bb *Backbuffer) Fill(color uint32) {

	b := byte(uint8(color))
	g := byte(uint8(color >> 8))
	r := byte(uint8(color >> 16))
	a := byte(uint8(color >> 24))

	for i := 0; i < len(bb.pixels); i += bb.bytesPerPixel {
		bb.pixels[i] = b
		bb.pixels[i+1] = g
		bb.pixels[i+2] = r
		bb.pixels[i+3] = a
	}
}
