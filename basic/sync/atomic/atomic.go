package main

import "sync/atomic"
func main(){
	atomic.AddInt32(*int32(1))
}
