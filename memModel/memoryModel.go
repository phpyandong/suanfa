package main

import (
	"sync"
	"fmt"
	"time"
)

var a string
var l sync.Mutex
var once sync.Once
var done bool
func f(){
	a = "hello world"
	l.Unlock()
}
func doPrint(){
	once.Do(setup)
	print(a)
}
func setup(){
	a = "hello ,steup"
	done = true
}
func main(){
	//l.Lock()
	//go f()
	//l.Lock()
	//println(a)
	//go doPrint()
	//go doPrint()
	//twoprint()
	go setup()
	for !done{
		fmt.Println(time.Now())
	}
	print(a)
}

func twoprint(){
	go doPrint()
	go doPrint()


}