package task12

import "fmt"

type Quant struct {
	field string
}

func makeQuant(s [5]string) []Quant {
	var r = make([]Quant, 0, 0)
	for _, v := range s {
		t := Quant{field: v}
		r = append(r, t)
	}
	return r
}

func RunTask() {
	s := [5]string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println("Последовательность строк:", s)
	fmt.Println("Множество:", makeQuant(s))
}
