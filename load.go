package images

import (
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path"
)

/*
图像加载
	1. 图片加载
	params:
		img_path: 图片文件路径
	returns:
		img: 图片对象
		err: error
*/
func Load(img_path string) (img image.Image, err error) {
	fi, err := os.Stat(img_path)
	if err != nil {
		return
	}
	if fi.IsDir() {
		err = errors.New("no file")
		return
	}
	fo, err := os.Open(img_path)
	if err != nil {
		return
	}
	defer fo.Close()
	switch path.Ext(img_path) {
	case ".jpg", ".jpeg":
		img, err = jpeg.Decode(fo)
		if err != nil {
			return
		}
	case ".png":
		img, err = png.Decode(fo)
		if err != nil {
			return
		}
	case ".gif":
		img, err = gif.Decode(fo)
		if err != nil {
			return
		}
	}
	return
}
