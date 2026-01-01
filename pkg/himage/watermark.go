package himage

import (
	"image"
	"image/color"

	"github.com/disintegration/imaging"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

// WatermarkText adds a text watermark to the image.
// This is a basic implementation using the basic font.
func (h *HImage) WatermarkText(text string, x, y int, col color.Color) *HImage {
	// We need a mutable image to draw on.
	// imaging.Overlay is cleaner if we create an image from text, but drawing text directly is also possible.
	// For simplicity and quality, let's create a label and overlay it.
	// However, drawing text in Go requires `freetype` or `font` packages and can be verbose.
	// Let's use a simple drawer on the image clone.

	// Create a copy of the image to draw on
	img := imaging.Clone(h.Image)

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  fixed.Point26_6{X: fixed.I(x), Y: fixed.I(y)},
	}
	d.DrawString(text)

	h.Image = img
	return h
}

// WatermarkImage adds an image watermark at the specified position.
func (h *HImage) WatermarkImage(watermarkPath string, x, y int, opacity float64) *HImage {
	if h.Err != nil {
		return h
	}
	wm, err := imaging.Open(watermarkPath)
	if err != nil {
		h.Err = err
		return h
	}

	h.Image = imaging.Overlay(h.Image, wm, image.Pt(x, y), opacity)
	return h
}

// WatermarkImageObj allows passing an image object directly
func (h *HImage) WatermarkImageObj(wm image.Image, x, y int, opacity float64) *HImage {
	h.Image = imaging.Overlay(h.Image, wm, image.Pt(x, y), opacity)
	return h
}
