package main

import (
	"fmt"
	"sync"
)
//处理工作
func doWorker(id int ,c chan int ,done chan bool){
	for n := range c{
		fmt.Printf("Worker %d received %c \n",id,n)
		go func() {
			done <- true
		}()
	}
}
//处理工作
func doWorker2(id int  ,w worker2){
	for n := range w.inChanel{
		fmt.Printf("Worker %d received %c \n",id,n)
		fmt.Println(w)
		w.waitGroupDone()
	}
}
type worker struct {
	inChanel chan int
	doneChanel chan bool
}
type worker2 struct {
	inChanel chan int
	waitGroupDone func()
}
//创建工作
func createWorker(id int) worker{
	w := worker{
		inChanel: make(chan int),
		doneChanel :make(chan bool),
	}
	go doWorker(id,w.inChanel,w.doneChanel)
	return w
}
//创建工作
func createWorker2(id int,group *sync.WaitGroup) worker2{

	w := worker2{
		inChanel: make(chan int),
		waitGroupDone :func(){
			group.Done()
		},
	}
	go doWorker2(id,w)
	return w
}
func chanDemo(){
	var workers [10] worker

	for i:= 0; i<10 ;i++  {

		workers[i] = createWorker(i)

	}

	for i:=0;i<10 ;i++  {

		workers[i].inChanel <- 'a'+i

		//<-workers[i].doneChanel

	}

	for i:=0;i<10 ;i++  {

		workers[i].inChanel <- 'A'+i

		//<-workers[i].doneChanel

	}

	for _,worker := range workers {

		fmt.Printf("%s",<-worker.doneChanel)

		fmt.Printf("%s",<-worker.doneChanel)

	}

}

func chanDemo2(){
	//创建200计数器
	var waitGroup sync.WaitGroup
	waitGroup.Add(200)
	//waitGroup 对象内部有计数器，
	// 最初从0开始，三个方法Add(计数器设置n),Done(计数器) ,Wait(阻塞代码的运行
	// 直到计数器减为0)
	// 用来控制计数器的数量


	var workers [100]worker2

	for i:= 0; i<100 ;i++  {
		//一定要通过指针传值，不然进程会进入死锁状态
		workers[i] = createWorker2(i,&waitGroup)

	}
	for i:=0;i<100 ;i++  {

		workers[i].inChanel <- 'a'
		//<-workers[i].doneChanel

	}

	for i:=0;i<100 ;i++  {

		workers[i].inChanel <- 'A'

		//<-workers[i].doneChanel

	}

	waitGroup.Wait()
}
func main(){
	chanDemo2()
}