package main

import (
	"sync"
	"fmt"
	"time"
)

var Wait sync.WaitGroup
var Counter int = 0
//bug 示范 go run -race waitgrp.go  WARNING: DATA RACE
func main()  {
	for routine :=1;routine <=2; routine++{
		Wait.Add(1)//计数加1 ，在方法前增加
		go Routine(routine)
	}
	Wait.Wait()//阻塞等待直到计数为0
	fmt.Println("final Count :%d \n ",Counter)

}
func Routine(id int ){
	for count:= 0; count < 2; count++{
		value := Counter
		time.Sleep(1 * time.Nanosecond)//协程触发调度，导致bug 高频复现
		value ++ //i++ 底层是 三行汇编 要去使用mutext
		Counter = value
	}
	Wait.Done()//该协程执行完毕后，调用done方法
}
