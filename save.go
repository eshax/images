package images

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path"
	"sync"
)

func Save(img_path string, img image.Image) (err error) {
	// log.Println(img_path)
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

/*
保存瓦片到文件

input:
 targetPath: 目标路径
      tiles: 瓦片数组
   tileSize: 瓦片尺寸
          x: 位于目标逻辑大图上的横向坐标
          y: 位于目标逻辑大图上的纵向坐标
*/
func SaveTiles(targetPath string, tiles [][]image.Image, tileSize, x, y int) (err error) {
	os.MkdirAll(targetPath, 0755)
	var wg sync.WaitGroup
	p := 0
	for c := range tiles {
		for r := range tiles[c] {
			wg.Add(1)
			p++
			go func(targetPath string, c, r int) {
				Save(fmt.Sprintf("%s/%d_%d.jpg", targetPath, (c+(x/tileSize)), (r+(y/tileSize))), tiles[c][r])
				wg.Done()
			}(targetPath, c, r)
			if p == ThreadPool {
				wg.Wait()
				p = 0
			}
		}
	}
	wg.Wait()
	return
}
