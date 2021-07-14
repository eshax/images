package images

import (
	"log"

	"github.com/disintegration/imaging"
)

/*
等比例调整图片大小, 原路径替换文件
	in:
		filepath: 图片文件路径, 包含文件名
		   scale: 缩放比例 (正数放大, 负数缩小)
	out:
		err
	example:
		images.Resize("image/a.jpg",   2) // 放大一倍
		images.Resize("image/a.jpg", 0.5) // 缩小一倍
*/
func Resize(filepath string, scale float64) error {

	// 缩放比例不能是 0
	if scale == 0 {
		scale = 1
	}

	// 原图的尺寸
	width, height, err := Info(filepath)
	if err != nil {
		return err
	}

	// 加载图像文件
	img, err := imaging.Open(filepath)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	width = int(float64(width) * scale)
	height = int(float64(height) * scale)

	// 图像尺寸不能小于 1
	if width == 0 {
		width = 1
	}
	if height == 0 {
		height = 1
	}

	img = imaging.Resize(img, width, height, imaging.Lanczos)
	err = imaging.Save(img, filepath)
	if err != nil {
		return err
	}

	return nil
}
