package images

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path"
)

func Save(img_path string, img image.Image) (err error) {
	o1, err := os.Create(img_path)
	if err != nil {
		return err
	}
	defer o1.Close()
	switch path.Ext(img_path) {
	case ".jpg", ".jpeg":
		if err := jpeg.Encode(o1, img, &jpeg.Options{Quality: 95}); err != nil {
			return err
		}
	case ".png":
		if err := png.Encode(o1, img); err != nil {
			return err
		}
	case ".gif":
		if err := gif.Encode(o1, img, &gif.Options{}); err != nil {
			return err
		}
	}
	return nil
}
