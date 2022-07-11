package task14

import "fmt"

func typeOfVar(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("int:", v)
	case string:
		fmt.Println("string:", v)
	case bool:
		fmt.Println("bool:", v)
	case chan int:
		fmt.Println("channel:", v)
	default:
		fmt.Println("unknown:", v)
	}
}

func RunTask() {
	a := 1
	b := "test"
	c := true
	d := make(chan int)
	typeOfVar(a)
	typeOfVar(b)
	typeOfVar(c)
	typeOfVar(d)
}
