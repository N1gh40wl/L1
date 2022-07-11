package task11

import (
	"fmt"
	"math/rand"
	"time"
)

func RunTask() {
	rand.Seed(time.Now().UnixNano())
	a := make([]int, 0, 0)
	b := make([]int, 0, 0)
	for i := 0; i < rand.Intn(15+1); i++ {
		a = append(a, rand.Intn(5)+1)
	}
	for i := 0; i < rand.Intn(15+1); i++ {
		b = append(b, rand.Intn(5)+1)
	}
	fmt.Println(a)
	fmt.Println()
	fmt.Println(b)
	fmt.Println()
	for _, av := range a {
		for _, bv := range b {
			if av == bv {
				println(bv)
			}
		}
	}
}
