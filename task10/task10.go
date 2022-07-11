package task10

import (
	"fmt"
	"strconv"
	"strings"
)

func RunTask() {
	s := "-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, -9.53, 8.45"
	fmt.Println(s)
	fmt.Println()
	temp := map[string][]float64{}
	split := strings.Split(s, ", ")

	for _, v := range split {
		t, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Println(err)
		}
		if t < 0 {
			temp["-"+strconv.Itoa((int(t)/10)*10*-1)] = append(temp["-"+strconv.Itoa((int(t)/10)*10*-1)], t)
		} else {
			temp[strconv.Itoa((int(t)/10)*10)] = append(temp[strconv.Itoa((int(t)/10)*10)], t)
		}

	}
	for k, v := range temp {
		fmt.Printf("group: %s, temp: %f\n", k, v)
	}
}
