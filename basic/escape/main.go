package main

/*
	go build -gcflags '-m -l' escape1.go
    go tool objdummmp -s "main\.main" test
	没有逃逸
	值传递 直接在栈上分配
	Go 语言函数传递都是通过值的，调用函数的时候，直接在栈上copy出一份参数
	不存在逃逸
 */
//type S1 struct {
//
//}
//func indenty1(x S1) S1 {
//	return x
//}
//
//type S2 struct {
//
//}
//x 未发生逃逸
//identity 函数的输入直接当成返回值了，因为没有对z作引用，
//所以z没有逃逸
//对x的引用也没有逃出main函数的作用域，因此x也没有发生逃逸
//func indenty2(z *S2) *S2{
//	return z
//}
//var x S2
//y := &x
//_ = *indenty2(y)

//
//type S3 struct {
//
//}

/**
	z 发生逃逸
z 是对x的拷贝，indenty3 函数中对z取了引用，所以z不能放在栈上。
否则在indenty3函数之外，通过引用不能找到z,所以z必须要逃逸到堆上
尽管在main函数中，直接丢弃了identy3的结果，但是Go的编译器，
还没有那么自能，分析不出来这种情况。
而对x从来就没有取引用，所以 x 不会发生逃逸
 */
//func indenty3(z S3) *S3{
//	return &z
//}
//var x S3
//_ = *indenty3(x)

//type S4 struct {
//	M *int
//}
///**
//	y 逃逸
//	refstruct 函数对y取了引用，所以y发生了逃逸
// */
//func refStruct4(y int)(z S4){
//	z.M = &y
//	return z
//}
//var i int
//refStruct4(i)

//type S5 struct {
//	M *int
//}
/**
	i 没有发生逃逸
在main函数里对i取了引用，并且把它传给了refstrutc 函数，i的引用
一直在main函数的作用域用，因此i没有发生逃逸
和上一个例子相比，i的写法有一点小差别，但是i的命运是不同的。
导致的程序效果是不同的，
例子4中，i先在main的栈帧中分配，
本例中，i只分配了一次，然后通过引用传递
 */
//./main.go:93:11: leaking param: y //参数y 溢出
//./main.go:93:19: ref6 z does not escape //参数z 没有逃逸
//./main.go:123:6: moved to heap: i //i逃逸到堆上


//func refStruct5(y *int)(z S5)  {
//	z.M = y
//	return z
//}
//var i int
//refStruct5(&i)
type S6 struct {
	M *int
}

//		i 逃逸，x未逃逸。x的作用域一直在main函数中
//	本例i发生了逃逸，按照前面例子5的分析，i不会逃逸，
//但是此例y赴给了一个输入的struct ,go 不能从输出返回到输入
//两个例子的区别是例子5中的S是在返回值里的，输入只能"流入"到输出
//本例中的S是在输入参数中，所以逃逸分析失败，i要逃逸到堆上

func ref6(y *int ,z *S6){
	z.M = y
}
//var x S6
//var i int
//ref6(&i,&x)



func main() {
	//var x S1
	//_ = indenty1(x)

	//var x S2
	//
	//y := &x
	//_ = *indenty2(y)
	//var x S3
	//_ = *indenty3(x)

	//var i int
	//refStruct4(i)

	//var i int
	//refStruct5(&i)
	//
	var x S6
	var i int
	ref6(&i,&x)
}