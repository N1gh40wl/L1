package task4

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func worker(ch <-chan interface{}, i int) {
	for v := range ch {
		time.Sleep(1 * time.Second)
		fmt.Println("Воркер номер: ", i+1, " Прочитал: ", v)

	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func RunTask() {
	rand.Seed(time.Now().UnixNano())
	var n int
	fmt.Println("Введите кол-во воркеров:")
	fmt.Scanf("%d", &n)
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	channelMain := make(chan interface{})
	for i := 0; i < n; i++ {
		go worker(channelMain, i)
	}

	for {
		s := RandomString(4)
		select {
		case channelMain <- s:
		case <-shutdown:
			fmt.Println("Конец работы программы")
			return
		}
	}

}
