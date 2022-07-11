package task5

import (
	"fmt"
	"time"
)

func reader(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}



func timer(ch chan<- int, n int) {

	time.Sleep(time.Second * time.Duration(n))
	ch <- 0
}

func RunTask() {
	var n int
	fmt.Println("Введите кол-во секунд:")
	fmt.Scanf("%d", &n)
	channel := make(chan int)
	channelTimer := make(chan int)
	go timer(channelTimer, n)
	go reader(channel)
	for i := 1; ; i++ {
		select {
		case <-channelTimer:
			return
		default:
			channel <- i
		}
	}
}
