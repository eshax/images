package images

import (
	"errors"
	"fmt"
	"image"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/disintegration/imaging"
)

/*
获取矩阵信息
	params:
		sourcePath: 源文件路径
	returns:
		cols: 列
		rows: 行
*/
func GetMatrixInfo(sourcePath string) (cols, rows, blockWidth, blockHeight, logicWidth, logicHeight int, extension string, err error) {
	fileList, e := ioutil.ReadDir(sourcePath)
	if e != nil {
		log.Println("GetMatrixInfo:ReadDir:", e)
		return
	}
	blockWidth, blockHeight, logicWidth, logicHeight = 0, 0, 0, 0
	for i := range fileList {
		if fileList[i].IsDir() {
			continue
		}
		name := fileList[i].Name()
		extension = path.Ext(name)
		name = strings.ReplaceAll(name, extension, "")
		data := strings.Split(name, "_")
		if len(data) != 2 {
			err = errors.New("filename is wrong")
			return
		}
		col, e := strconv.Atoi(data[0])
		if e != nil {
			log.Println("GetMatrixInfo:strconv.Atoi:", e)
			return
		}
		row, e := strconv.Atoi(data[1])
		if e != nil {
			log.Println("GetMatrixInfo:strconv.Atoi:", e)
			return
		}
		cols = Max(cols, col+1)
		rows = Max(rows, row+1)

		if blockWidth == 0 || blockHeight == 0 {
			img, e := Load(path.Join(sourcePath, fileList[i].Name()))
			if e != nil {
				log.Println("GetMatrixInfo:Load:", e)
				return
			}
			blockWidth = img.Bounds().Max.X
			blockHeight = img.Bounds().Max.Y
		}
	}
	logicWidth = blockWidth * cols
	logicHeight = blockHeight * rows
	return
}

func bufferToTiles(sourcePath, targetPath, extension string, maxLevel, minLevel, logicWidth, logicHeight, blockWidth, blockHeight, bufSize, tileSize int) {

	os.MkdirAll(targetPath, 0755)

	if ((logicWidth % 2) > 0) || ((logicHeight % 2) > 0) {
		bufSize = Max(bufSize, Max(logicWidth, logicHeight))
	}

	left, top := 0, 0
	for left < logicWidth {
		for top < logicHeight {
			img, _ := Buffer(sourcePath, extension, bufSize, logicWidth, logicHeight, blockWidth, blockHeight, left, top)
			if sourcePath != fmt.Sprintf("%s/%d", targetPath, maxLevel) {
				SaveTiles(fmt.Sprintf("%s/%d", targetPath, maxLevel), Tile(img, tileSize), tileSize, left, top)
			}
			halfBufferToTiles(img, targetPath, maxLevel, minLevel, tileSize, left/2, top/2)
			top += bufSize
		}
		left += bufSize
		top = 0
	}

}

func halfBufferToTiles(img image.Image, targetPath string, maxLevel, minLevel, tileSize, left, top int) {
	log.Println(maxLevel, fmt.Sprintf("(%d, %d) - (%d, %d)", left, top, img.Bounds().Max.X+left, img.Bounds().Max.Y+top))
	if maxLevel > minLevel {
		img = imaging.Resize(img, (img.Bounds().Max.X / 2), (img.Bounds().Max.Y / 2), imaging.Lanczos)
		SaveTiles(fmt.Sprintf("%s/%d", targetPath, maxLevel-1), Tile(img, tileSize), tileSize, left, top)
		halfBufferToTiles(img, targetPath, maxLevel-1, minLevel, tileSize, left/2, top/2)
	}
}

func getBufferLevels(logicWidth, logicHeight, bufSize, tileSize int) (int, int) {

	bufSize = BufferSize(bufSize, tileSize)

	max := GetMaxLevel(logicWidth, logicHeight)

	if logicWidth <= bufSize && logicHeight <= bufSize {
		return max, max
	}

	bufSizeCount := 0
	_bufSize := bufSize
	for {
		if _bufSize <= tileSize {
			break
		}
		_bufSize /= 2
		bufSizeCount++
	}

	oddCount := 0
	for {
		if (((logicWidth % bufSize) % 2) > 0) || (((logicHeight % bufSize) % 2) > 0) {
			break
		}
		if logicWidth <= bufSize && logicHeight <= bufSize {
			break
		}
		logicWidth /= 2
		logicHeight /= 2
		oddCount++
	}

	min := Max((max - bufSizeCount), (max - oddCount))

	return max, min
}

func Matrix(sourcePath, targetPath string, bufSize, tileSize int) {
	cols, rows, blockWidth, blockHeight, logicWidth, logicHeight, extension, err := GetMatrixInfo(sourcePath)

	log.Println("LogicWidth:", logicWidth, "LogicHeight:", logicHeight, "Cols:", cols, "Rows:", rows, "BlockWidth:", blockWidth, "BlockHeight:", blockHeight)

	if err != nil {
		log.Println("GetMatrixInfo:", err)
		return
	}

	bufSize = BufferSize(bufSize, tileSize)
	log.Println("Buffer Size:", bufSize)
	log.Println("Tile Size:", tileSize)

	for {
		max, min := getBufferLevels(logicWidth, logicHeight, bufSize, tileSize)
		if max == min {
			bufferToTiles(sourcePath, targetPath, extension, min, 1, logicWidth, logicHeight, blockWidth, blockHeight, bufSize, tileSize)
			break
		}
		bufferToTiles(sourcePath, targetPath, extension, max, min, logicWidth, logicHeight, blockWidth, blockHeight, bufSize, tileSize)
		logicWidth, logicHeight = GetLevelSize(logicWidth, logicHeight, min)
		blockWidth, blockHeight = tileSize, tileSize
		sourcePath = fmt.Sprintf("%s/%d", targetPath, min)
	}

}
