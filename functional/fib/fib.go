package main

import (
	"fmt"
	"strings"
	"io"
	"bufio"
)
type intGen func() int


// 1, 1, 2, 3, 5, 8, 13, ...
//		 	a  b

func Fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
func fib() func()int{
	a , b := 0,1
	return func() int {
		a,b = b, a+b
		return a
	}
}
func febinacci2()func() int{
	a ,b := 0,1
	return func() int {
		a,b = b ,a+b
		return a
	}
}
//不停读取feibo
//func (g intGen) Read(
//	//继承
//	p []byte) (n int,err error){
//		next := g()
//	if next > 1000 {
//		return 0,io.EOF
//	}
//	fmt.Println("我是read 继承",p)
//	s := fmt.Sprintf("%d\n",next)
//	return strings.NewReader(s).Read(p)
//}


func (g intGen) Read(p []byte) (n int, err error){
	next := g()
	if  next > 1000{

	}
	s := fmt.Sprintf("%d\n",next)
	return strings.NewReader(s).Read(p)

}
func printFileContents(reader io.Reader){
	scanner := bufio.NewScanner(reader)
	for scanner.Scan()  {
		fmt.Println("哈",scanner.Text())
	}
}
func main() {
	f := Fibonacci()
	printFileContents(f)




}