package window

import (
	"github.com/Ganners/gorender/backbuffer"
	"github.com/veandco/go-sdl2/sdl"
)

// Creates and returns a Window
func NewWindow(appName string, width, height int) *Window {
	return &Window{
		WindowName: appName,
		Width:      width,
		Height:     height,
	}
}

// Window is an SDL window which can be used to draw bitmaps onto
type Window struct {
	Window     *sdl.Window
	surface    *sdl.Surface
	WindowName string
	Width      int
	Height     int
	Backbuffer *backbuffer.Backbuffer
}

// Creates a window with the configuration of it's receiver
func (w *Window) Create() {

	// Initialize and create the window
	sdl.Init(sdl.INIT_EVERYTHING)
	window, err := sdl.CreateWindow(
		w.WindowName,
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		w.Width,
		w.Height,
		sdl.WINDOW_SHOWN)

	if err != nil {
		panic(err)
	}

	// If successful, set the window
	w.Window = window

	// Grab the window surface
	w.surface, err = window.GetSurface()
	if err != nil {
		panic(err)
	}

	// Create and attach a backbuffer to work with
	w.Backbuffer = backbuffer.New(
		w.Width,
		w.Height,
		w.surface.BytesPerPixel(),
		w.surface.Pixels())

	// Fill background black and update
	w.Backbuffer.Fill(0xFF000000)
	w.Update()
}

// Updates the surface, should be called if backbuffer has been updated
func (w *Window) Update() {
	w.Window.UpdateSurface()
}

// Destroy should be executed when we are finished running
func (w *Window) Destroy() {
	w.Window.Destroy()
}

// Returns the width and height
func (w *Window) GetSize() (int, int) {
	return w.GetSize()
}

// Sets the width and height
func (w *Window) SetSize(width, height int) {
	w.Window.SetSize(width, height)
}

// Minimizes the window
func (w *Window) Minimize() {
	w.Window.Minimize()
}

// Maximizes the window
func (w *Window) Maximize() {
	w.Window.Maximize()
}

// Hides the window
func (w *Window) Hide() {
	w.Window.Hide()
}

// Shows the window
func (w *Window) Show() {
	w.Window.Show()
}

// Sets the Window to fullscreen
func (w *Window) SetFullscreen() {
	w.Window.SetFullscreen(sdl.WINDOW_FULLSCREEN)
}
