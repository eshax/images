package images

import (
	"log"
	"os"
	"testing"
)

func Test_Matrix(t *testing.T) {
	os.RemoveAll("dist/lower/dzi")
	Matrix("dist/lower", "dist/lower/dzi", 4000, 256)
}

func Test_GetLevelSize(t *testing.T) {
	// GetLevelSize(16384, 19584, 15, 15)
	// GetLevelSize(16384, 19584, 14, 15)
	// GetLevelSize(16384, 19584, 13, 15)
	// GetLevelSize(16384, 19584, 12, 15)
	// GetLevelSize(16384, 19584, 11, 15)
	// GetLevelSize(16384, 19584, 10, 15)
	// GetLevelSize(16384, 19584, 9, 15)
	// GetLevelSize(16384, 19584, 8, 15)
	// GetLevelSize(16384, 19584, 7, 15)
	// GetLevelSize(16384, 19584, 6, 15)
	// GetLevelSize(16384, 19584, 5, 15)
	// GetLevelSize(16384, 19584, 4, 15)
	// GetLevelSize(16384, 19584, 3, 15)
	// GetLevelSize(16384, 19584, 2, 15)
	// GetLevelSize(16384, 19584, 1, 15)
}

func Test_GetBufferLevels(t *testing.T) {
	log.Println(getBufferLevels(163840, 195840, 10240, 256))
}
