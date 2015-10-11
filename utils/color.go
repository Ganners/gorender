package utils

func ColorToRGBABytes(color uint32) (r, g, b, a byte) {

	b = byte(uint8(color))
	g = byte(uint8(color >> 8))
	r = byte(uint8(color >> 16))
	a = byte(uint8(color >> 24))

	return
}
