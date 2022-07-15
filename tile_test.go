package images

import (
	"log"
	"testing"
)

func Test_GetTileRange(t *testing.T) {
	log.Println(getTileRange(100, 100, 256))
	log.Println(getTileRange(256, 256, 256))
	log.Println(getTileRange(500, 500, 256))
	log.Println(getTileRange(1024, 1024, 256))
}

func Test_GetTileRect(t *testing.T) {
	log.Println(GetTileRect(1024, 1024, 1, 1, 256))
}
