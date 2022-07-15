package images

import (
	"errors"
	"fmt"
	"image"
	"image/draw"
	"log"
	"sync"
)

func SquareMerge(source_folder_path, extension string, logicWidth, logicHeight, blockSize int) (img image.Image, err error) {

	if logicWidth%blockSize > 0 || logicHeight%blockSize > 0 {
		return nil, errors.New("logic width or logic height is wrong")
	}

	cols := logicWidth / blockSize
	rows := logicHeight / blockSize

	dst := image.NewRGBA(image.Rect(0, 0, logicWidth, logicHeight))

	var wg sync.WaitGroup
	pool := 0
	for col := 0; col < cols; col++ {
		for row := 0; row < rows; row++ {
			go func(source_folder_path, extension string, blockSize, col, row int) {

				img, err := Load(fmt.Sprintf("%s/%d_%d%s", source_folder_path, col, row, extension))
				if err != nil {
					log.Println(err)
				}
				bounds := img.Bounds()
				if blockSize != bounds.Max.Y || blockSize != bounds.Max.X {
					log.Println(errors.New("block width or block height is wrong"))
				}
				l, t := blockSize*col, blockSize*row
				r, b := blockSize+l, blockSize+t
				draw.Draw(dst, image.Rect(l, t, r, b), img, image.Point{}, draw.Over)

			}(source_folder_path, extension, blockSize, col, row)

			if pool == ThreadPool {
				wg.Wait()
				pool = 0
			}
		}
	}
	wg.Wait()

	return dst, nil
}
