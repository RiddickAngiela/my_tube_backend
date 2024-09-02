package utils

import (
	"bytes"
	"image"
	"image/jpeg"
	"io"
	"os"

	"golang.org/x/image/draw"
)

// ResizeImage resizes an image to the specified width and height.
func ResizeImage(file io.Reader, width, height int) error {
	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	dst := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.CatmullRom.Scale(dst, dst.Bounds(), img, img.Bounds(), draw.Over, nil)

	var buf bytes.Buffer
	err = jpeg.Encode(&buf, dst, nil)
	if err != nil {
		return err
	}

	err = os.WriteFile("./uploads/resized_image.jpg", buf.Bytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}
