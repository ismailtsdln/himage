# himage - Advanced Image Processing CLI

`himage` is a powerful and flexible command-line tool for image processing written in Go. It allows you to resize, convert, filter, and watermark images using a beautiful terminal user interface. It supports both single-file and batch processing.

## Features

- **Batch Processing**: Recursively process entire directories of images.
- **Resize**: Smart resizing with aspect ratio preservation.
- **Convert**: Convert between common image formats (JPEG, PNG, GIF, BMP, TIFF).
- **Filters**: Apply various filters including Sepia, Sigmoid, Blur, Sharpen, and more.
- **Watermark**: Add text or image watermarks with customizable opacity and positioning.
- **Info**: Inspect image metadata.
- **Modern UI**: Interactive terminal output using `pterm`.

## Installation

### Go Install

You can install `himage` directly using the `go install` command:

```bash
go install -v github.com/ismailtsdln/himage/cmd/himage@latest
```

### From Source

```bash
git clone https://github.com/ismailtsdln/himage.git
cd himage
go build -o himage ./cmd/himage
```

## Usage

### Batch Processing

All commands (`resize`, `convert`, `filter`, `watermark`) support directory inputs.

```bash
# Resize all images in 'photos' directory
./himage resize photos/ --width 800 --output resized_photos/
```

### Info

Get information about an image.

```bash
./himage info input.jpg
```

### Resize

Resize an image to specific dimensions.

```bash
./himage resize input.jpg --width 800 --height 600 --output resized.jpg
```

### Convert

Convert an image from one format to another.

```bash
./himage convert input.png output.jpg
```

### Filters

Apply filters to your images.

**Sepia (Retro Effect)**

```bash
./himage filter input.jpg --type sepia --output vintage.jpg
```

**Sigmoid (Contrast)**

```bash
./himage filter input.jpg --type sigmoid --value 5.0 --output contrast.jpg
```

**Available Filters:**

- `sepia`
- `sigmoid` (value: contrast factor)
- `grayscale`
- `blur` (requires value)
- `sharpen` (requires value)
- `invert`
- `brightness` (value: -100 to 100)
- `contrast` (value: -100 to 100)
- `gamma` (value > 0)
- `saturation` (value: -100 to 100)

### Watermark

Add a text or image watermark.

**Text Watermark**

```bash
./himage watermark input.jpg --text "Copyright 2024" --x 10 --y 10 --output watermarked.jpg
```

**Image Watermark**

```bash
./himage watermark input.jpg --image logo.png --x 10 --y 10 --opacity 0.5 --output watermarked.jpg
```

## License

MIT License
