package himage

import (
	"image"
	"image/color"
	"testing"
)

func createTestImage(w, h int) *HImage {
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	// Fill with some color
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			img.Set(x, y, color.RGBA{uint8(x % 255), uint8(y % 255), 100, 255})
		}
	}
	return &HImage{
		Image: img,
		Path:  "test_image.png",
		Ext:   ".png",
	}
}

func TestResize(t *testing.T) {
	img := createTestImage(100, 100)
	img.Resize(50, 50)

	bounds := img.Image.Bounds()
	if bounds.Dx() != 50 || bounds.Dy() != 50 {
		t.Errorf("Expected size 50x50, got %dx%d", bounds.Dx(), bounds.Dy())
	}
}

func TestResizeAspectRatio(t *testing.T) {
	img := createTestImage(100, 50)
	// Resize width to 50, keep aspect ratio (height should become 25)
	img.Resize(50, 0)

	bounds := img.Image.Bounds()
	if bounds.Dx() != 50 || bounds.Dy() != 25 {
		t.Errorf("Expected size 50x25, got %dx%d", bounds.Dx(), bounds.Dy())
	}
}

func TestRotate(t *testing.T) {
	img := createTestImage(100, 50)
	img.Rotate90()

	bounds := img.Image.Bounds()
	// 90 deg rotation swaps dimensions
	if bounds.Dx() != 50 || bounds.Dy() != 100 {
		t.Errorf("Expected size 50x100 after 90deg rotation, got %dx%d", bounds.Dx(), bounds.Dy())
	}
}

func TestFilters(t *testing.T) {
	img := createTestImage(100, 100)

	// Test Grayscale
	img.Grayscale()
	// Hard to inspect pixels easily without lengthy code, but we ensure it doesn't panic

	// Test Blur
	img.Blur(2.0)

	// Test Invert
	img.Invert()

	// Test Sepia
	img.Sepia()

	if img.Image == nil {
		t.Error("Image became nil after filters")
	}
}

func TestWatermarkText(t *testing.T) {
	img := createTestImage(100, 100)
	img.WatermarkText("Test", 10, 10, color.White)

	if img.Image == nil {
		t.Error("Image became nil after watermark")
	}
}
