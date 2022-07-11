package task21

import (
	"fmt"
	"strconv"
	"strings"
)

func showStats(s string) {
	vals := strings.Fields(s)
	w, _ := strconv.ParseFloat(vals[0], 32)
	h, _ := strconv.ParseFloat(vals[1], 32)
	a, _ := strconv.Atoi(vals[2])
	fmt.Printf("Вес: %f фунтов\nРост: %f футов\nВозраст: %d лет\n", w, h, a)
}

func RunTask() {
	person := &person{}
	usa := &usa{
		weight: 194.6,
		height: 5.7,
		age:    22,
	}
	person.inserStatsInFunc(usa)
	fmt.Println()
	russian := &russia{
		weight: 87.5,
		height: 181.3,
		age:    20,
	}
	russiaAdapter := &russiaAdapter{
		russianHuman: russian,
	}
	person.inserStatsInFunc(russiaAdapter)
}
