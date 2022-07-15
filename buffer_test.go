package images

import (
	"fmt"
	"os"
	"testing"
)

func Test_GetBufferRange(t *testing.T) {
	logicWidth := 16384
	logicHeight := 19584
	blockWidth := 256
	blockHeight := 256
	left := 0
	top := 0
	bufSize := 40000
	getBufferRange(logicWidth, logicHeight, blockWidth, blockHeight, left, top, bufSize)
}

func Test_Buffer(t *testing.T) {
	bufSize := 2048
	tileSize := 256
	logicWidth := 2048 * 8
	logicHeight := 2448 * 8
	blockWidth := 2048
	blockHeight := 2448
	left := bufSize * 7
	top := bufSize * 9
	// scale := 0.5
	buf, _ := Buffer("dist/lower", ".jpg", bufSize, logicWidth, logicHeight, blockWidth, blockHeight, left, top)
	buf = ResizeImage(buf, 0.5)
	Save(fmt.Sprintf("dist/lower.join.%d.%d.%d.%d.%d.%d.%d.jpg", bufSize, logicWidth, logicHeight, blockWidth, blockHeight, left, top), buf)
	tiles := Tile(buf, tileSize)
	os.RemoveAll("dist/dzi")
	SaveTiles("dist/dzi", tiles, tileSize, left/2, top/2)
}
