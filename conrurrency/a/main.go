package main
var c = make(chan int, 10)

var a string
func f() {
	a = "hello, world"
	//c <- 0
	close(c)//关闭或传一个数据 ，channal 就不会阻塞
}

func main1() {
	go f()
	<-c  //阻塞
	print(a)
}

var limit = make(chan int, 3)

func main() {
	for _, w := range work {
		go func(w func()) {
			limit <- 1
			w()
			<-limit
		}(w)
	}
	select{}
}