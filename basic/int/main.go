package main

import (
	"fmt"
	"unsafe"
)
type myFuncType func(int,int) int


func main() {
	var c byte = 255//0-255
	var a int8 = 127//-127  127
	var b uint8  = 255//0-255
	var n1 = 1100
	fmt.Printf("n1类型%T n2 占用字节数是%d\n",1,unsafe.Sizeof(n1))
	fmt.Println(c,a,b)
	var d int = '北'
	fmt.Printf("这个比较吊了,按字符输出 :%c\n",d)
	a1 :=1
	a2 :=2
	a1 ,a2 = a2 ,a1
	fmt.Println(a1,a2)
	fmt.Println(6/1*3)
	arr2 := []int{1,123}
	testArr(arr2)
	fmt.Println(arr2)

	fmt.Printf("%T\n",arr2)
	var str string = "xx"
	testName(str)
	fmt.Println(str)
	fmt.Println(MyFunc(getSum,1,2))

	res := func (n1 ,n2 int) int{
		return n1+n2
	}(10,20)
	fmt.Println("res:",res)

	res2 :=func (n1,n2 int) int{
		return n1+n2
	}
	res3:=res2(10,233)
	fmt.Println("res3:",res3)
}
func MyFunc(funct myFuncType ,num1 int,num2 int) int{
	return funct(num1,num2)

}
func getSum(num1 ,num2 int ) int{
	return num2 +num1
}

func testName(str string) {
	str ="哈哈"

}
func testArr(arr []int)  {
	arr[1]++
}
