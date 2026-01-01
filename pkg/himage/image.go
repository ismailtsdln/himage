package himage

import (
	"image"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
)

// HImage wraps the standard image.Image to provide chainable methods.
type HImage struct {
	Image image.Image
	Path  string
	Ext   string
	Err   error
}

// Error returns the accumulated error, if any.
func (h *HImage) Error() error {
	return h.Err
}

// Load loads an image from the specified file path.
func Load(path string) (*HImage, error) {
	img, err := imaging.Open(path)
	if err != nil {
		return nil, err
	}

	return &HImage{
		Image: img,
		Path:  path,
		Ext:   strings.ToLower(filepath.Ext(path)),
	}, nil
}

// Save saves the image to the original path.
func (h *HImage) Save() error {
	if h.Err != nil {
		return h.Err
	}
	// If path is empty, we can't save to original
	if h.Path == "" {
		return imaging.ErrUnsupportedFormat
	}
	return imaging.Save(h.Image, h.Path)
}

// SaveAs saves the image to the specified path.
func (h *HImage) SaveAs(path string) error {
	if h.Err != nil {
		return h.Err
	}
	return imaging.Save(h.Image, path)
}

// SaveQuality saves the image with the specified quality (mostly for JPEG).
func (h *HImage) SaveQuality(path string, quality int) error {
	return imaging.Save(h.Image, path, imaging.JPEGQuality(quality))
}
