package images

import (
	"log"
	"testing"
)

func TestInfo(t *testing.T) {
	w, h, err := Info("dist/0_0.jpg")
	if err != nil {
		log.Println(err.Error())
		t.Error(err)
	}
	t.Log(w, h)
}

func TestResize(t *testing.T) {
	if err := Resize("dist/0_0.jpg", 2); err != nil {
		t.Error(err)
	}
}
