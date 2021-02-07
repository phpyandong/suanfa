package main

import (
	"fmt"
	"time"
	"sync"
)

type atomicInt struct {
	value int
	lock sync.Mutex
}

func (a *atomicInt) increment(){
	fmt.Println("safe increment ")//执行多次，不加锁
	func(){
		//代码区块内加锁，利用匿名函数。只执行一次，由于加锁
		a.lock.Lock()
		defer  a.lock.Unlock()
		a.value++
	}()
}
func (a *atomicInt) get() int{
	a.lock.Lock()
	defer a.lock.Unlock()

	return  a.value
}

func main(){
	var a atomicInt
	a.increment()
	go func() {
		a.increment()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a.get())

}
