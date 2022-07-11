package task2

import "fmt"

func load(a [5]int, c chan<- int) {
	for _, v := range a {
		c <- v
	}
	close(c)
}

func sqrt(a chan int, c chan<- int) {
	for v := range a {
		c <- (v * v)
	}
	close(c)
}

func RunTask() {
	a := [5]int{2, 4, 6, 7, 10}
	val := make(chan int)
	res := make(chan int)
	go load(a, val)
	go sqrt(val, res)

	for v := range res {
		fmt.Println(v)
	}

}
