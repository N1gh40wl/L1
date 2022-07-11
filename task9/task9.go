package task9

import (
	"sync"
)

func reader(arr [5]int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(ch)
	for _, v := range arr {
		ch <- v
	}
}

func square(arr <-chan int, arrSq chan<- int, wg *sync.WaitGroup) {
	for {
		foo, err := <-arr
		if !err {
			close(arrSq)
			wg.Done()
			return
		}
		arrSq <- (foo * 2)
	}

}

func printer(ch <-chan int, wg *sync.WaitGroup) {
	for {
		foo, err := <-ch
		if !err {
			wg.Done()
			return
		}
		println(foo)
	}
}

func RunTask() {

	arr := [5]int{1, 2, 3, 4, 5}
	arrChan := make(chan int)
	arrSqChan := make(chan int)
	var wg sync.WaitGroup
	wg.Add(3)
	go reader(arr, arrChan, &wg)
	go square(arrChan, arrSqChan, &wg)
	go printer(arrSqChan, &wg)

	wg.Wait()

}
