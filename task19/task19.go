package task19

import (
	"fmt"
	"strings"
)

func ReverseWords(str string) string {
	s := strings.Fields(str)
	var res []string

	for _, e := range s {
		var r string
		for i := len(e) - 1; i >= 0; i-- {
			r = r + string(e[i])
		}
		res = append(res, r)
	}

	return strings.Join(res, " ")
}

func RunTask() {
	var s string
	fmt.Println("Введите слово:")
	fmt.Scanf("%s\n", &s)
	unicode := []rune(s)
	var res string
	for i := len(unicode) - 1; i >= 0; i-- {
		res = res + string(unicode[i])
	}
	fmt.Println(res)
}
