package task21

import (
	"fmt"
	"strconv"
)

type russia struct {
	weight float32
	height float32
	age    int
}

func (r russia) conertToUSA() string {
	return fmt.Sprintf("%f", r.weight*2.2) + " " + fmt.Sprintf("%f", r.height*0.032) + " " + strconv.Itoa(r.age)
}
