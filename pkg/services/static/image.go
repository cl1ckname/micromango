package static

import (
	"bytes"
	"fmt"
	"golang.org/x/image/draw"
	"image"
	"image/jpeg"
	"image/png"
	pb "micromango/pkg/grpc/static"
	"os"
)

// ToImage converts a PNG image to JPEG format
func ToImage(imageBytes []byte, from pb.ImageType) (image.Image, error) {
	switch from {
	case pb.ImageType_PNG:
		// Decode the PNG image bytes
		return png.Decode(bytes.NewReader(imageBytes))
	case pb.ImageType_JPG:
		return jpeg.Decode(bytes.NewReader(imageBytes))
	}

	return nil, fmt.Errorf("unable to convert %#v to jpeg", from)
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
