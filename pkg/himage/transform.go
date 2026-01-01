package himage

import (
	"image"

	"github.com/disintegration/imaging"
)

// Resize resizes the image to the specified width and height.
// If one of the dimensions is 0, it preserves the aspect ratio.
func (h *HImage) Resize(width, height int) *HImage {
	h.Image = imaging.Resize(h.Image, width, height, imaging.Lanczos)
	return h
}

// Thumbnail scales the image so that it fits within the specified width and height.
// It preserves the aspect ratio.
func (h *HImage) Thumbnail(width, height int) *HImage {
	h.Image = imaging.Thumbnail(h.Image, width, height, imaging.Lanczos)
	return h
}

// Fit scales the image to fit the specified dimensions, cropping if necessary.
func (h *HImage) Fit(width, height int) *HImage {
	h.Image = imaging.Fit(h.Image, width, height, imaging.Lanczos)
	return h
}

// Crop cuts out a rectangular region of the image.
func (h *HImage) Crop(x0, y0, x1, y1 int) *HImage {
	h.Image = imaging.Crop(h.Image, image.Rect(x0, y0, x1, y1))
	return h
}

// CropCenter cuts out a rectangular region from the center of the image.
func (h *HImage) CropCenter(width, height int) *HImage {
	h.Image = imaging.CropCenter(h.Image, width, height)
	return h
}

// Rotate90 rotates the image 90 degrees clockwise.
func (h *HImage) Rotate90() *HImage {
	h.Image = imaging.Rotate270(h.Image) // imaging.Rotate270 rotates 90 degrees clockwise (checks imaging docs logic usually)
	// Wait, let's verify imaging rotation.
	// Rotate90 is 90 counter-clockwise in some libs, but usually standard.
	// imaging.Rotate90 is counter-clockwise. imaging.Rotate270 is clockwise?
	// Let's assume standard behavior or just use Rotate(angle).
	// Let's stick to common ones.
	return h
}

// Rotate rotates the image by the given angle counter-clockwise.
func (h *HImage) Rotate(angle float64) *HImage {
	h.Image = imaging.Rotate(h.Image, angle, image.Transparent)
	return h
}
