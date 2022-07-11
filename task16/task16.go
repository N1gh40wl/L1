package task16

import (
	"fmt"
	"sort"
)

type test struct {
	name string
	val  int
}

type ByVal []test

func (a ByVal) Len() int           { return len(a) }
func (a ByVal) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByVal) Less(i, j int) bool { return a[i].val < a[j].val }

func RunTask() {
	vals := []int{6, 53, 33, 7, 3, 21, 4}
	fmt.Println("Массив до сортировки:   ", vals)
	sort.Ints(vals)
	fmt.Println("Массив после сортировки:", vals)
	fmt.Println()

	testVals := []test{
		{"first", 4},
		{"second", 1},
		{"third", 2},
	}
	fmt.Println(testVals)
	fmt.Println()

	sort.Sort(ByVal(testVals))
	fmt.Println(testVals)
	sort.Slice(testVals, func(i, j int) bool {
		return testVals[i].val > testVals[j].val
	})
	fmt.Println()
	fmt.Println(testVals)
}
