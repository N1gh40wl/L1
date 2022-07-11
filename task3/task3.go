package task3

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

func sum(a chan int, c chan<- int) {
	var res int
	for v := range a {
		res += v
	}
	c <- res
	close(c)
}

func RunTask() {
	a := [5]int{2, 4, 6, 7, 10}
	val := make(chan int)
	sqrtVals := make(chan int)
	res := make(chan int)
	go load(a, val)
	go sqrt(val, sqrtVals)
	go sum(sqrtVals, res)

	for v := range res {
		fmt.Println(v)
	}

}
