package images

import (
	"fmt"
	"image"
	"image/draw"
	"log"
	"sync"
)

/*
将瓦片文件合并成一张大图
*/
func MergeFiles(source_path string) image.Image {

	cols, rows, blockWidth, blockHeight, logicWidth, logicHeight, extension, err := GetMatrixInfo(source_path)

	if err != nil {
		log.Println("iamges.MergeFiles:", err)
		return &image.NRGBA{}
	}

	dst := image.NewRGBA(image.Rect(0, 0, logicWidth, logicHeight))

	var wg sync.WaitGroup
	wg.Add(cols * rows)
	for col := 0; col < cols; col++ {
		for row := 0; row < rows; row++ {
			go func(source_path, extension string, bw, bh, col, row int) {
				defer wg.Done()
				img, err := Load(fmt.Sprintf("%s/%d_%d%s", source_path, col, row, extension))
				if err != nil {
					log.Println("iamges.MergeFiles:", err)
				}
				bounds := img.Bounds()
				l, t := blockWidth*col, blockHeight*row
				r, b := bounds.Max.X+l, bounds.Max.Y+t
				draw.Draw(dst, image.Rect(l, t, r, b), img, image.Point{}, draw.Over)
			}(source_path, extension, blockWidth, blockHeight, col, row)
		}
	}
	wg.Wait()

	return dst
}
