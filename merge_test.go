package images

import (
	"log"
	"testing"
)

func Test_MergeFiles(t *testing.T) {

	log.Println()

	img := MergeFiles("dist/15")
	Save("dist/15.merge.jpg", img)

	log.Println()
}
