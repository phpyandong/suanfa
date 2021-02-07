package main

import (
	"fmt"
	"time"
	"math/rand"
)
func generator() chan int{
	out := make(chan int)
	go func() {
		i:=0
		for {
			time.Sleep(
				time.Duration(rand.Intn(1500)) *time.Millisecond,
			)
			out <- i
			i++
		}
	}()
	return out
}
func worker(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second)//消费太慢的话，数据会被覆盖，应该加本地队列
		fmt.Printf("Worker %d received %d\n",
			id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}
func main()  {
	var c1,c2 chan int = generator(),generator()
	var worker = createWorker(0) //chan int //nil := createWorker(0)
	n := 0
	//hasValue := false
	var values []int //实现本地队列，防止数据挤压，数组
	tm := time.After(100* time.Second)//定时器，10秒钟后，从这个channel得到时间
	tick := time.Tick(1 * time.Second)
	for {
		var activeWorker chan<- int
		//if hasValue {
		//	activeWorker = worker
		//}
		var activeValue int
		if len(values) >0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select{
		case n = <-c1:
			//hasValue = true
			values = append(values,n)
		case n = <-c2:
			//hasValue = true
			values = append(values,n)
		case activeWorker <- activeValue:
			//hasValue = false
			values = values[1:]//删除第一个元素。
		case <-time.After(1* time.Second):
			fmt.Println("timeOut")//打不出来
		case c:=<- tm://从这个定时器channel 获取时间
			fmt.Println("bye",c)
			return
		case <-tick:
			fmt.Println("length:", len(values))
		default:

		}
	}

}

func main1()  {
	var c1,c2 chan int = generator(),generator()
	var worker = createWorker(0) //chan int //nil := createWorker(0)
	n := 0
	hasValue := false

	for {
		var activeWorker chan<- int
		if hasValue {
			activeWorker = worker
		}
		select{
			case n = <-c1:
				hasValue = true
			case n = <-c2:
				hasValue = true
			case activeWorker <- n:
				//hasValue = false
			default:

		}
	}

}

func main2()  {
	var c1,c2 chan int = generator(),generator()
	for {
		select{
		case n:= <-c1:
			fmt.Println("receive form c1:",n)
		case n:= <-c2:
			fmt.Println("reveive from c2:",n)
		default:
		}
	}

}