package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type atomicInt struct {
	value int
	lock  sync.Mutex
}

func (i *atomicInt) increment() {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.value++
}

func (i *atomicInt) get() int {
	i.lock.Lock()
	defer i.lock.Unlock()
	return i.value
}

func panicTest() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		panic(err)
	}
	defer func() {
		resp.Body.Close()
		fmt.Println("jhjjjjjj")
	}()
	fmt.Println(resp)
}

func main() {
	aInt := &atomicInt{}
	aInt.increment()
	for i := 0; i < 5; i++ {
		go func() {
			aInt.increment()
		}()
	}
	time.Sleep(time.Second)
	fmt.Println("result = ", aInt.value)
	panicTest()
}
