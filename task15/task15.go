package task15

import (
	"fmt"
	"math/rand"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func createHugeString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

//var justString string //глобальная переменная не очень хороший инструмент для разработки

func someFunc(justString *string) {
	v := createHugeString(1 << 10) //не была объявлена функция
	if len(v) >= 100 {
		*justString = v[:100] //мы не знаем что нам вернула createHugeString может быть ошибка
	}
	fmt.Println("v:", v)
	fmt.Println()
	fmt.Println("j:", *justString)
}

func RunTask() {
	justString := ""
	rand.Seed(time.Now().UnixNano())
	someFunc(&justString)
}
