package task25

import (
	"fmt"
	"time"
)

func timer(t int) {
	<-time.After(time.Second * time.Duration(t))
}

func RunTask() {
	fmt.Println("Введите время сна:")
	var t int
	fmt.Scanf("%d",&t)
	timer(t)
}
