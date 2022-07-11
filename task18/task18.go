package task18

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	num   int
	mutex sync.Mutex
}

func (c *Counter) Inc(i int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.num += 1
	fmt.Println(c.num, i)
	time.Sleep(time.Millisecond * 120)
}

func timer(ch chan<- int) {
	time.Sleep(time.Second * 1)
	ch <- 1
}

func RunTask() {
	c := Counter{num: 0}
	end := make(chan int)
	go timer(end)
	for i := 0; ; i++ {
		select {
		case <-end:
			fmt.Println("Счетчик:", c.num)
			return
		default:
			go c.Inc(i)
		}
	}
}
