package images

import (
	"image"
	"os"
)

func Info(path string) (width, height int, err error) {

	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()

	img, _, err := image.DecodeConfig(f)
	if err != nil {
		return
	}

	width = img.Width
	height = img.Height

	return
}
