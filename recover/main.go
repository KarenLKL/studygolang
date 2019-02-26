package main

import "fmt"

func main() {
	tryRecover()
}

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Printf("it is a error:%v", err)
		} else {
			panic(fmt.Sprintf("i do not know what to do:%v", r))
		}
	}()

	//a, b := 0, 5
	//fmt.Println(b / a)
	panic("what to do?")
}
