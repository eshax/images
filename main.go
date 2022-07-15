package images

import (
	"math"
)

var ThreadPool = 100

func Min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func Max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func GetMaxLevel(width, height int) (max_level int) {
	return int(math.Ceil(math.Log2(math.Max(float64(height), float64(width)))))
}

func BufferSize(bufSize, tileSize int) int {
	return int(math.Pow(2, float64(int(math.Log2(float64(bufSize))))))
}

func GetLevelSize(logicWidth, logicHeight, level int) (int, int) {
	levels := GetMaxLevel(logicWidth, logicHeight)
	if level > levels {
		return 0, 0
	}
	for i := levels; i > level; i-- {
		logicWidth /= 2
		logicHeight /= 2
	}
	return logicWidth, logicHeight
}
