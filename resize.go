package images

import (
	"log"
	"os"
	"path"

	"github.com/disintegration/imaging"
)

/*
等比例调整图片大小, 原路径替换文件
	in:
		source_path: 源文件路径
		target_path: 目标文件路径
		scale: 缩放比例 (正数放大, 负数缩小)
	out:
		err
	example:
		images.Resize("image/0/a.jpg", "image/1/b.jpg",   2) // 放大一倍
		images.Resize("image/0/a.jpg", "image/1/b.jpg", 0.5) // 缩小一倍
*/
func Resize(source_path, target_path string, scale float64) error {

	target_dir := path.Dir(target_path)

	if err := os.MkdirAll(target_dir, 0755); err != nil {
		return err
	}

	// 缩放比例不能是 0
	if scale == 0 {
		scale = 1
	}

	// 原图的尺寸
	width, height, err := Info(source_path)
	if err != nil {
		return err
	}

	// 加载图像文件
	img, err := imaging.Open(source_path)
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
	err = imaging.Save(img, target_path)
	if err != nil {
		return err
	}

	return nil
}
