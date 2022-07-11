package task13

import "fmt"

func RunTask() {
	a := 1
	b := 2
	fmt.Println("Числа до смены a:", a," b:",b)
	a, b = b ,a 
	fmt.Println("Числа после смены a:", a," b:",b)
}
