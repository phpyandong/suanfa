package main

import "fmt"

func lean(){
	a := 1243
	fmt.Printf("%v\n",a)

	fmt.Printf("a的指针%v\n",&a)

	var point *int = &a
	fmt.Printf("point指针：%v\n",&point)
	fmt.Printf("point内存的a的指针：%v\n",point)
	fmt.Printf("point内存的a的值：%v\n",*point)

}
func demo(){
	var a int = 300
	var b int = 800

	fmt.Println("a 的指针",&a)
	//var ptr *int = a 错误的写法
	var ptr *int = &a
	*ptr = 200
	fmt.Printf("修改后的值%d\n",a)

	ptr2 := &b
	*ptr2 = 880
	fmt.Println("修改后b的值",b)

	var kh string = "22"
	fmt.Println("kh",&kh)
	f := &kh

	fmt.Println("f ,",f)

}

func main(){
	demo()




}
