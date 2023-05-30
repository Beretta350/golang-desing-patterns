package adapter

import (
	"fmt"
)

func minmax(a, b int) (int, int) {
	if a < b {
		return a, b
	} else {
		return b, a
	}
}

func Main() {
	rc := NewRectangle(6, 4)
	a := VectorToRaster(rc)
	_ = VectorToRaster(rc)
	fmt.Print(DrawPoints(a))
}
