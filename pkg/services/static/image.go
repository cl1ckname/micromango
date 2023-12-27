package static

import (
	"bytes"
	"fmt"
	"golang.org/x/image/draw"
	"golang.org/x/image/tiff"
	"golang.org/x/image/webp"
	"image"
	"image/jpeg"
	"image/png"
	"micromango/pkg/grpc/share"
	"os"
	"strings"
)

// ToImage converts an image to JPEG format
func ToImage(image *share.File) (image.Image, error) {
	ext, err := getMime(image.Filename)
	if err != nil {
		return nil, err
	}
	switch ext {
	case "png":
		return png.Decode(bytes.NewReader(image.File))
	case "jpg", "jpeg":
		return jpeg.Decode(bytes.NewReader(image.File))
	case "webp":
		return webp.Decode(bytes.NewReader(image.File))
	case "tiff":
		return tiff.Decode(bytes.NewReader(image.File))
	default:
		return nil, fmt.Errorf("unable to convert %#v to jpeg: invalid mime %s", image.Filename, ext)
	}
}

func Resize(src image.Image, w, h int) image.Image {
	targetRect := image.Rect(0, 0, w, h)
	dst := image.NewRGBA(targetRect)
	// Resize:
	draw.NearestNeighbor.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)
	return dst
}

func SaveImage(path string, img image.Image) error {
	buf := new(bytes.Buffer)

	// encode the image as a JPEG file
	if err := jpeg.Encode(buf, img, nil); err != nil {
		return err
	}
	imageBytes := buf.Bytes()
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	if _, err := f.Write(imageBytes); err != nil {
		return err
	}
	return nil
}

func getMime(filename string) (string, error) {
	filename = strings.ToLower(filename)
	nameParts := strings.Split(filename, ".")
	if len(nameParts) != 2 {
		return "", fmt.Errorf("invalid filename: %s (len is %d)", filename, len(nameParts))
	}
	return nameParts[1], nil
}
