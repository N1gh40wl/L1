package task6

import (
	"fmt"
	"sync"
	"time"
)

func reader(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

func producer(ch chan<- int, stop <-chan int) { // 1 - Остановка с помощью доп. канала
	for i := 0; ; i++ {
		select {
		case <-stop:
			return
		default:
			ch <- i

		}
	}

}

func stopByPointer(stop *bool) { //2 - Остановка через указатель
	for i := 0; ; i++ {
		if *stop {
			fmt.Println(i)
		} else {
			return
		}
	}
}

func timer(ch chan<- int, n int) {
	time.Sleep(time.Duration(n) * time.Second)
	ch <- 0
}

func PointerTimer(stop *bool, n int) {
	time.Sleep(time.Duration(n) * time.Second)
	*stop = true
}

func stopByWaitGroup(ch <-chan int, wg *sync.WaitGroup) { // 3 - Остановка через закрытие канала
	for {
		foo, err := <-ch
		if !err {
			wg.Done()
			return
		}
		println(foo)
	}
}

func timerWG(ch chan<- int, n int) {
	newTimer := time.NewTimer(time.Second * time.Duration(n))
	for i := 0; ; i++ {
		select {
		case <-newTimer.C:
			return
		default:
			ch <- i
		}
	}
}

func RunTask() {

	var n int
	fmt.Println("Введите кол-во секунд:")
	fmt.Scanf("%d", &n)

	//	1 - Остановка с помощью доп. канала
	channel := make(chan int)
	timeChan := make(chan int)
	go timer(timeChan, n)
	go reader(channel)
	producer(channel, timeChan)

	/*	2 - Остановка через указатель
		stop := true
		go stopByPointer(&stop)
		PointerTimer(&stop, n)
	*/

	/*	3 - Остановка через закрытие канала
		var wg sync.WaitGroup
		wg.Add(1)
		go stopByWaitGroup(channel, &wg)
		timerWG(channel, n)
		close(channel)
		wg.Wait()
	*/
}
