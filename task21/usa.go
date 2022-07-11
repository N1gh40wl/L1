package task21

import (
	"fmt"
	"strconv"
)

type usa struct {
	weight float32
	height float32
	age    int
}

func (h *usa) insertHumanStats() {
	showStats(fmt.Sprintf("%f", h.weight) + " " + fmt.Sprintf("%f", h.height) + " " + strconv.Itoa(h.age))
}
