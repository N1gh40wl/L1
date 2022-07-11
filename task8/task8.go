package task8

import (
	"fmt"
	"strconv"
	"strings"
)

func iterativeDigitsCount(number int64) int {
	count := 0
	for number != 0 {
		number /= 10
		count += 1
	}
	return count

}

func changeString(s string, i int) string {
	c := s[i : i+1]
	if c == "0" {
		c = "1"
	} else {
		c = "0"
	}
	res := s[:i] + c + s[i+1:]
	return res
}



func RunTask() {
	var val int64
	var n int
	fmt.Println("Введите число: ")
	fmt.Scanf("%d", &val)
	fmt.Println("Введите номер бита:")
	fmt.Scanf("%d", &n)
	fmt.Println("Число до преобразования:   ", val, ", в двоичном представлении:",
		strconv.FormatInt(val, 2))
	fmt.Println(strings.Repeat(" ", 27+iterativeDigitsCount(val)+30+n) + "|")
	valAfter, _ := strconv.ParseInt(changeString(strconv.FormatInt(val, 2), n), 2, 64)
	fmt.Println("Число после преобразования:", valAfter, ", в двоичном представлении:",
		strconv.FormatInt(valAfter, 2))
}
