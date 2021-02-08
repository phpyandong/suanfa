package copy_on_write

import (
	"sync/atomic"
	"time"
	"sync"
)

func loadConifg() string{
	return "test"
}
func expence0(){
	type Map map[string]string
	var m atomic.Value
	m.Store(make(Map))
	var mu sync.Mutex
	read := func(key string) (val string) {
		m1 := m.Load().(Map)
		return m1[key]
	}
	insert := func(Key,val string) {
		mu.Lock()
		defer mu.Unlock()
		m1 := m.Load().(Map)
		m2 := make(Map) //写时复制
		for k , v := range m1{
			m2[k] = v
		}
		m2["key"] = val
		m.Store(m2) //原子替换，场景十分类似
					// 新建文件 + 修改别名
					//新建es索引 + 修改别名
	}
	_,_ = read,insert
}

//实例1 伪代码 将config安全的拉取远程更新到本地
func expcence1(){
	var config atomic.Value
	config.Store(loadConifg())
	config.Store(loadConifg())
	go func() {
		for {
			time.Sleep( 10* time.Second)
			config.Store(loadConifg())
		}
	}()
	requests := []string{"aaa","bbb"}
	for i:= 0;i<10 ;i++  {
		go func() {
			for r := range requests{
				c := config.Load()
				_,_ = r,c

			}
		}()
	}
}