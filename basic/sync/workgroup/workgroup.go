package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println(context.Background())
	tr := NewTracker()
	//把并发丢给调用者。我们使用服务端埋点来记录一些事件
	//无法保证创建的goroutine 声明周期管理，
	// 会导致最常见的问题，就是在服务关闭的时候，有一些事件丢失
	go tr.Run()//

	_ = tr.Event(context.Background(),"text")
	_ = tr.Event(context.Background(),"text")
	_ = tr.Event(context.Background(),"text")
	//把控何时推出，context 定时推出
	ctx,cancel := context.WithDeadline(context.Background(),time.Now().Add(5*time.Second))
	defer cancel()
	tr.Shutdown(ctx)
}

func NewTracker() *Tracker{
	return &Tracker{
		ch :make(chan string ,10),
	}
}
//从channel去取消息，消费消息做数据上报
func (t *Tracker) Run(){
	for data :=range t.ch {
		time.Sleep(1 * time.Second)
		fmt.Printf(data)//上传到Elk等日志平台
	}
	t.stopChanl <- struct{}{}//t.ch.close() 后，给stop信号
}
type Tracker struct {
	ch chan string
	stopChanl chan struct{}
}
//处理消息，发送到channel
func (t *Tracker) Event(ctx context.Context,data string) error{
	select {
		case t.ch <- data://将数据发送到channel,暂存10条
			return nil
	case <- ctx.Done()://??????
		return ctx.Err()
	}
}
func (t *Tracker) Shutdown(ctx context.Context){
	close(t.ch) //控制Run()方法的退出，即上报日志的goroutine退出
	select {
		case <-t.stopChanl://得到 主动关闭的信号
		case <-ctx.Done():
	}
}