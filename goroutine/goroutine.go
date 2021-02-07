package main

import (
	"time"
	"fmt"
	"runtime"
)

func main() {
	//for i := 0; i < 1000; i++ {
	//	go func(i int) {
	//		for {
	//			fmt.Printf("Hello from "+
	//				"goroutine %d\n", i)
	//		}
	//	}(i)
	//}
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			//for j := i; j>0 ;j-- {
			for {
				a[i]++
				runtime.Gosched()//主动让出控制权，否则会死循环，
			}
			//}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}