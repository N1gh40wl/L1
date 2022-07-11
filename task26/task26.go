package task26

import (
	"fmt"
	"sort"
	"strings"
)

type SortBy []rune

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i] < a[j] }

func RunTask() {
	s := ""
	fmt.Println("Введите слово:")
	fmt.Scanf("%s\n", &s)
	fmt.Printf("%s - ", s)
	s = strings.ToLower(s)
	unicode := []rune(s)
	sort.Sort(SortBy(unicode))
	f := true
	for i := 0; i < len(unicode)-1; i++ {
		if unicode[i] == unicode[i+1] {
			f = false
		}
	}

	fmt.Printf("%t\n", f)
}
