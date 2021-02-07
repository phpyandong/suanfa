package main

import "fmt"

//闭包

func A(i int){
	i++
	fmt.Println(i)
}

func create()(func()int)  {
	c := 2 //c为捕获变量
	return func() int {
		return c
	}
}
func create2()(fs [2]func()){
	for i:=0;i<2 ;i++  {
		fs[i] = func() {
			fmt.Println(i)
		}
	}
	//fmt.Println("create2:",fs)
	return
}
func test2()  {
	fs := create2()
	for i:=0;i<len(fs) ;i++  {
		fs[i]()
	}
}
func test1()  {
	f1 := create()
	f2 := create()
	fmt.Println(f1())
	fmt.Println(f2())

}
var Name string = "tom"
//Name := "tom"//这种方式会报错 因为 为 var  name  string ;name="tom" 赋值时报错
func main() {
	str :="ws我这"
	s:= []rune(str)
	fmt.Println(len(s))
	fmt.Println(Name)

	test1()
	test2()

}
func B(){
	f1 :=A
	f1(1)
}
func C(){
	f2 := A
	f2(1)
}


