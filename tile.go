package images

import (
	"image"
	"sync"

	"github.com/disintegration/imaging"
)

/*
计算瓦片的数量, 行、列各有多少个
*/
func getTileRange(width, height, tileSize int) (cols, rows int) {

	cols = width / tileSize
	rows = height / tileSize

	if (width % tileSize) > 0 {
		cols += 1
	}

	if (height % tileSize) > 0 {
		rows += 1
	}

	return
}

/*
从原始逻辑图像中提取瓦片数据

input:
	     src: 原始图像
	tileSize: 瓦片尺寸

output:
	     dst: 瓦片数组
*/
func Tile(src image.Image, tileSize int) (dst [][]image.Image) {
	w := src.Bounds().Max.X
	h := src.Bounds().Max.Y
	// log.Println("images.Tile:", "width:", w, "height:", h)
	cols, rows := getTileRange(w, h, tileSize)
	dst = make([][]image.Image, cols)
	for c := range dst {
		dst[c] = make([]image.Image, rows)
	}
	var wg sync.WaitGroup
	p := 0
	for c := 0; c < cols; c++ {
		for r := 0; r < rows; r++ {
			wg.Add(1)
			p += 1
			go func(c, r int) {
				defer wg.Done()
				dst[c][r] = imaging.Crop(src, image.Rect(GetTileRect(w, h, c, r, tileSize)))
			}(c, r)
			if p == ThreadPool {
				wg.Wait()
				p = 0
			}
		}
	}
	wg.Wait()
	return
}

/*
计算瓦片在原始逻辑图像中的坐标
params:
	 width: 逻辑图宽度
	height: 逻辑图高度
	   col: 列 (瓦片)
	   row: 行 (瓦片)
  tileSize: 瓦片尺寸 (正方形)
*/
func GetTileRect(width, height, col, row, tileSize int) (left, top, right, bottom int) {

	left = col * tileSize
	top = row * tileSize

	right = Min((left + tileSize), width)
	bottom = Min((top + tileSize), height)

	return
}
