package main

import (
	"fmt"
	"time"
)

func channelDemo(){
	//var c chan int     //c == nil
	c := make(chan int)
	c <- 1
	c <- 2
	n := <-c
	fmt.Println(n)
}
func worker(c chan int){
	func(){
		for{
			n := <-c
			fmt.Printf("worke %d rev %d",n)
		}

	}()
}
func worke2(woker_id int ,c chan int){
	func(){
		for{
			n := <-c
			println(fmt.Sprintf("worker %d recv %c",woker_id,n))
		}

	}()
}
func createWorker(id int) chan<- int {
	c := make(chan int)
	go func() {
		for {
			fmt.Println(fmt.Sprintf("carete worker %d receive %c \n",id,<-c))
		}
	}()
	return c
}
func channelDemo2(){
	//var c chan int     //c == nil
	c := make(chan int)
	//go func(){
	//	for{
	//		n := <-c
	//		fmt.Println(n)
	//	}
	//
	//}()
	go worker(c)
	c <- 1
	c <- 2
	time.Sleep(time.Millisecond)
	}
func channelDemo3(){
	//var c chan int     //c == nil
	var channels [10]chan<- int
	//go func(){
	//	for{
	//		n := <-c
	//		fmt.Println(n)
	//	}
	//
	//}()
	for i:=0; i<10 ;i++  {
		channels[i] = createWorker(i)
		//go worke2(i,channels[i])

	}
	for i:=0; i<10 ;i++  {
		channels[i] <- 'a' +i
	}

	time.Sleep(time.Millisecond)
}
func worker2(id int , c chan  int){
	//for  {
	//	n , ok := <-c
	//	if !ok {
	//		break
	//	}
	//	fmt.Printf("worker %d rec %d \n",id,n)
	//}

	for n := range c{

		fmt.Printf("worker %d rec %d \n",id,n)
	}
}
func bufferChannel(){
	c := make(chan int,3)
	go worke2(0,c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'

}
func channelClose(){
	c := make(chan int)
	go worker2(0,c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	//close(c)
	time.Sleep(time.Millisecond)
}
func main()  {
	//channelDemo()//fatal error: all goroutines are asleep - deadlock!  需要使用goroutines 分开接收和发送
	//channelDemo3()
	//bufferChannel()
	channelClose()
	//buffer channle
	//close channel
}