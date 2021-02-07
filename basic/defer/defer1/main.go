package main

import "fmt"

//go build -0 test main.go
//go tool objdump -s "mian\.main" test

//func f1()(r int){
//	t := 5
// 1 赋值指令
//	r = t
// 2 defer 被插入到赋值与返回之间执行，这个例子中返回值r没被修改过
//	func(){
//		t = t+5
//	}
//	return
//}
//
//func f1() (r int){
//	t := 5
//	defer func() {
//		t = t + 5 //这里是因为 defer 是拷贝
//		fmt.Println("defer",t)
//
//	}()
//	t = 5+2
//	return t
//}

////func f2()(r int){
////	1.赋值
////	r = 1
////	//2 这里改的是r是之前传值传进去的r ,不会改变要返回的那个r值
////	func (r int){
////		r = r+5
////	}(r)
////	return
////}
//func f2()(r int){
//	//fmt.Printf("f2,r的地址%v 值为%d\n",&r,r)
//
//	defer func(r int){
//		r = r + 5
//		//fmt.Printf("f2,defer r的地址%v 值为%d\n",&r,r)
//
//	}(r)
//	return 2
//}
////func f3()(r int){
////	1 赋值
////	r = 1
////2 .这里改的是r 是之前传引用传进去的r ,会改变要返回的那个值
////	func (r *int){
////		*r = *r +5
////	}(&r)
////	//3 空的return
////	return
////}
//func f3() (r int){
//	//fmt.Println(r) //先赋值，后执行defer
//	defer func(r *int){
//		//fmt.Println("defer R:",*r)
//		*r = *r +5
//	}(&r)
//	return 1
//}
//func e1(){
//	var err error
//	defer fmt.Println("e1",err)
//}
//func e2(){
//	var err error
//	defer func() {
//		fmt.Println(err)
//	}()
//	err = errors.New("defer2 error")
//}
//func e3(){
//	var err error
//	defer func(err error) {
//		fmt.Println(err)
//	}(err)
//	err = errors.New("defer3 error")
//}
//
func main(){
	//fmt.Println(f1())//5
	//f2()//1
	//fmt.Println(f3())//6  回头研究闭包
	//e1()
	//e2()
	//e3()
	//testdefer5()
	testAcc()
	//testdefer6()
	//testdefer7()
	//fmt.Println(defer8())
	//defer9()
}
func testdefer5(){
	var a = Accumulator()
	fmt.Printf("%d\n",a(1))
	fmt.Printf("%d\n",a(10))
	fmt.Printf("%d\n",a(100))
	fmt.Println("_____________________")

	var b = Accumulator()
	fmt.Printf("%d\n",b(2))
	fmt.Printf("%d\n",b(10))
	fmt.Printf("%d\n",b(100))

}
func testAcc(){
	a := Acc()
	fmt.Println(a(11))
	fmt.Println(a(22))

}
func Acc()func(int) int{
	var x int

	return func(de int) int {
		x+= de
		return x
	}
}
//func testdefer6(){
//	defer fmt.Println("defer main")
//	var user = ""
//	go func() {
//		defer func() {
//			fmt.Println("defer caller")
//			if err := recover();err != nil {
//				fmt.Println("recover success .err:",err)
//
//			}
//		}()
//
//		func(){
//			defer func() {
//				fmt.Println("defer here")
//			}()
//
//
//			if user =="" {
//				fmt.Println("should set user ")
//				panic("should set user env.")
//			}
//			fmt.Println("after panic")
//		}()
//	}()
//	time.Sleep(1000)
//	fmt.Println("end of main fucntion")
//}
//
////闭包引用了X变量 ，a,b 可看作2个不同的实例，实例之间互不影响，实例内部
// x变量 是同一个地址，因此具有累加效应
func Accumulator()func(int) int{
	var x int
	return func(delta int) int {
		fmt.Printf("(%+v,%+v) - ",&x,x)
		x+= delta
		return x
	}
}

//func testdefer7(){
//	for i:=0;i<5;i++{
//		defer fmt.Println(i,1)
//	}
//	for i:=0;i<5 ;i++  {
//		defer func() {
//			fmt.Println(i ,2)
//		}()
//	}
//	for i:=0; i<5 ; i++  {
//		defer func() {
//			j := i
//			//fmt.Println(j,3)
//			fmt.Println(j,&j,i,&i,3)
//		}()
//	}
//	for i :=0 ;i<5 ;i++  {
//		j := i
//		defer fmt.Println(j,4)
//	}
//	for i:=0;i<5 ;i++  {
//		j := i
//		defer func() {
//			fmt.Println(j,5)
//		}()
//	}
//
//	for i:=0; i<5;i++  {
//		defer func(j int) {
//			fmt.Println(j,6)
//		}(i)
//	}
//}
//func defer8()( c ,d int){
//	a,b := 1,2
//	defer func(b int) {
//		a = a+b
//		//fmt.Println("defer:",a,b) //5,2
//	}(b)
//	//fmt.Println("第1次打印",a,b)
//
//	a = a+b
//	//fmt.Println("第二次打印",a,b)//3,2
//	return a ,b
//}
//func B(i int){
//	i = 9
//}
//func defer9(){
//	for i:=0;i<9 ;i++  {
//		defer B(i)
//	}
//
//}