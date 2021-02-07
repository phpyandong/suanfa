package main

import (
	"net/http"
	"context"
	"fmt"
)

func main(){
	done := make(chan error ,2)
	stop := make(chan struct{})
	go func() {
		done <- serve("localhost",nil,stop)
	}()
	go func() {
		done <- serve("localhost",nil,stop)
	}()
	var stopped bool
	for i := 0;i < cap(done);i++{
		if err := <-done;err != nil {
			fmt.Println("error:%v",err)
		}
		if !stopped{
			stopped = true
			close(stop)
		}

	}

}
func serve(addr string ,handler http.Handler,stop <-chan struct{}) error{
	s := http.Server{
		Addr: addr,
		Handler: handler,
	}
	go func() {
		<-stop //等待停止信号
		s.Shutdown(context.Background())
	}()
	return s.ListenAndServe()
}
