package images

import (
	"fmt"
	"log"
	"sync"
	"testing"
)

func Test_Single(t *testing.T) {

	log.Println()

	count := 10

	var wg sync.WaitGroup
	wg.Add(count)

	for i := 0; i < count; i++ {
		go func(i int) {
			defer wg.Done()
			Single("dist/preview.jpg", fmt.Sprintf("dist/preview/dzi.%d", i), 256)
		}(i)
	}

	wg.Wait()

	log.Println()

}
