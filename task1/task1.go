package task1

import "fmt"

type Human struct {
	name       string
	lastname   string
	patronymic string
}

type Action struct {
	Human
}

func (h *Human) HumanFunc(name string, lastname string, patronymic string) {
	h.name = name
	h.lastname = lastname
	h.patronymic = patronymic
}

func RunTask() {
	var a Action
	a.HumanFunc("Тест", "Тестов", "Тестович")
	fmt.Println("Задание 1")
	fmt.Println(a)
}
