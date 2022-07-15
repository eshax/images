package images

import (
	"errors"
	"fmt"
	"image"
	"image/draw"
	"sync"
)

func getBufferRange(logicWidth, logicHeight, blockWidth, blockHeight, left, top, bufSize int) (min_col, min_row, max_col, max_row, x, y int) {

	right := bufSize + left
	bottom := bufSize + top

	right = Min(right, logicWidth)
	bottom = Min(bottom, logicHeight)

	// 涉及到的文件索引 ( 最小 - 最大 )  x_y.jpg
	min_col = left / blockWidth
	min_row = top / blockHeight
	max_col = right / blockWidth
	max_row = bottom / blockHeight

	// 偏移量
	x = left % blockWidth
	y = top % blockHeight

	if logicWidth <= (blockWidth * max_col) {
		max_col -= 1
	}

	if logicHeight <= (blockHeight * max_row) {
		max_row -= 1
	}

	// log.Println("left:", left, "top:", top, "right:", right, "bottom:", bottom, "min_col:", min_col, "min_row:", min_row, "max_col:", max_col, "max_row:", max_row, " x:", x, "y:", y)

	return
}

/*
在逻辑大图中抠一个虚拟区域

在瓦片化的多个图像文件形成的逻辑大图中, 锁定一个固定大小的虚拟区域, 进行图像提取, 提取到内存中

      source_folder_path: 原图像文件夹地址
               extension: 源文件的扩展名
                 bufSize: 虚拟区域的固定大小
blockWidth & blockHeight: 原瓦片文件的宽度与高度
         startX & startY: 虚拟区域位于逻辑大图的左上角

*/
func Buffer(source_path, extension string, bufSize, logicWidth, logicHeight, blockWidth, blockHeight, left, top int) (buffer image.Image, err error) {

	if left < 0 || top < 0 {
		err = errors.New("left & top can't be less than 0")
		return
	}

	c1, r1, c2, r2, x, y := getBufferRange(logicWidth, logicHeight, blockWidth, blockHeight, left, top, bufSize)

	width, height := Min(logicWidth-left, bufSize), Min((logicHeight-top), bufSize)

	// log.Println("Buffer:", "\t", left, "\t", top, "\t", width, "\t", height, "\t", c1, "\t", r1, "\t", c2, "\t", r2, "\t", x, "\t", y)

	buf := image.NewRGBA(image.Rect(0, 0, width, height))

	var wg sync.WaitGroup
	p := 0

	for col := c1; col <= c2; col++ {
		for row := r1; row <= r2; row++ {

			wg.Add(1)
			p += 1

			go func(c, r int) {

				defer wg.Done()

				img, err := Load(fmt.Sprintf("%s/%d_%d%s", source_path, c, r, extension))
				if err != nil {
					return
				}

				iw, ih := img.Bounds().Max.X, img.Bounds().Max.Y
				left := (blockWidth * (c - c1)) - x
				top := (blockHeight * (r - r1)) - y
				right := (left + iw)
				bottom := (top + ih)

				draw.Draw(buf, image.Rect(left, top, right, bottom), img, image.Point{}, draw.Src)

			}(col, row)

			if p == ThreadPool {
				wg.Wait()
				p = 0
			}
		}
	}

	wg.Wait()

	return buf, nil
}
