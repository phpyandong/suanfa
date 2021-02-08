package main

import (
	"testing"
	"sync"
	"sync/atomic"
)
type Config struct {
	a []int
}
//go test -bench=. config_test.go
//对比两个方法的执行时间
//BenchmarkAtomic-4       218987641                5.58 ns/op
//BenchmarkMutex-4         1000000              3015 ns/op


func (c *Config) T(){}
func BenchmarkAtomic(b *testing.B){
	var v atomic.Value
	v.Store(&Config{})
	go func() {
		i := 0
		for {
			i++
			cfg := &Config{a:[]int{i,i+1,i+2,i+3,i+4}}
			v.Store(cfg)
		}
	}()
	var wg sync.WaitGroup
	for n:=0;n<4 ;n++  {
		wg.Add(1)
		go func() {
			for n:=0;n<b.N ;n++  {
				cfg := v.Load().(*Config)
				cfg.T()
				//fmt.Printf("%v\n",cfg)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
func BenchmarkMutex(b *testing.B){
	var l sync.RWMutex
	var cfg *Config

	go func() {
		i :=0
		for{
			i++
			l.Lock()
			cfg = &Config{
				a:[]int{
					i,i+1,i+2,i+3,i+4,i+5,
			}}
			l.Unlock()
		}
	}()
	var wg sync.WaitGroup

	for n:=0;n<4;n++{
		wg.Add(1)
		for n:=0 ;n< b.N ; n++  {
			l.RLock()
			cfg.T()
			l.RUnlock()
		}
		wg.Done()
	}
	wg.Wait()
}
