package  main

import (
	"fmt"
)

//go build -gcflags '-m -l' escape.go
//go tool compile -S escape_analysis.go
type User struct {
	ID     int64
	Name   string
	Avatar string
}
//执行了 runtime.newobject 方法，也就是确实是分配到了堆上。
//这是为什么呢？这是因为 GetUserInfo() 返回的是指针对象，引用被返回到了方法之外了。
//因此编译器会把该对象分配到堆上，而不是栈上。
// 否则方法结束之后，局部变量就被回收了，\
//岂不是翻车。所以最终分配到堆上是理所当然的。


func GetUserInfo() *User {
	return &User{
		ID: 666666,
		Name: "sim lou",
		Avatar: "https://www.baidu.com/avatar/666666",
	}
}
//看，该对象分配到栈上了。很核心的一点就是它有没有被作用域之外所引用，
// 而这里作用域仍然保留在 main 中，因此它没有发生逃逸。
func PrintStr() {
	str := new(string)

	*str = "hello world"
	//通过对其分析，可得知当形参为 interface 类型时，
	// 在编译阶段编译器无法确定其具体的类型。因此会产生逃逸，最终分配到堆上。


	//如果你有兴趣追源码的话，可以看下内部的 reflect.TypeOf(arg).Kind() 语句，
	// 其会造成堆逃逸，而表象就是 interface 类型会导致该对象分配到堆上。
	//reflect.TypeOf(str).Kind()
	fmt.Println(str)

}


func main()  {
	//u := GetUserInfo()
	//println(u.Name)
	PrintStr()
}
