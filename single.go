package images

import "fmt"

func Single(sourceFilePath, targetPath string, tileSize int) (err error) {

	sourceImage, err := Load(sourceFilePath)
	if err != nil {
		return err
	}

	maxLevel := GetMaxLevel(sourceImage.Bounds().Max.X, sourceImage.Bounds().Max.Y)

	for level := maxLevel; level > 0; level-- {
		tileImages := Tile(sourceImage, tileSize)
		SaveTiles(fmt.Sprintf("%s/%d", targetPath, level), tileImages, tileSize, 0, 0)
		sourceImage = ResizeImage(sourceImage, 0.5)
	}

	return
}
