package main

import (
	"L1/task1"
	"L1/task10"
	"L1/task11"
	"L1/task12"
	"L1/task13"
	"L1/task14"
	"L1/task15"
	"L1/task16"
	"L1/task17"
	"L1/task18"
	"L1/task19"
	"L1/task2"
	"L1/task20"
	"L1/task21"
	"L1/task22"
	"L1/task23"
	"L1/task24"
	"L1/task25"
	"L1/task26"
	"L1/task3"
	"L1/task4"
	"L1/task5"
	"L1/task6"
	"L1/task7"
	"L1/task8"
	"L1/task9"
	"fmt"
)

func runTasks(i int) {
	switch i {
	case 1:
		task1.RunTask()
	case 2:
		task2.RunTask()
	case 3:
		task3.RunTask()
	case 4:
		task4.RunTask()
	case 5:
		task5.RunTask()
	case 6:
		task6.RunTask()
	case 7:
		task7.RunTask()
	case 8:
		task8.RunTask()
	case 9:
		task9.RunTask()
	case 10:
		task10.RunTask()
	case 11:
		task11.RunTask()
	case 12:
		task12.RunTask()
	case 13:
		task13.RunTask()
	case 14:
		task14.RunTask()
	case 15:
		task15.RunTask()
	case 16:
		task16.RunTask()
	case 17:
		task17.RunTask()
	case 18:
		task18.RunTask()
	case 19:
		task19.RunTask()
	case 20:
		task20.RunTask()
	case 21:
		task21.RunTask()
	case 22:
		task22.RunTask()
	case 23:
		task23.RunTask()
	case 24:
		task24.RunTask()
	case 25:
		task25.RunTask()
	case 26:
		task26.RunTask()
	}

}

func main() {
	for {
		var n int
		fmt.Scanf("%d\n", &n)
		runTasks(n)
		fmt.Println()
	}

}
