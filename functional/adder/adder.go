package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}
func adder3() func(int) int{
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		// 第一个加， 下一个要加
		return base + v, adder2(base + v)
	}
}




func main() {
	b := adder3()

	fmt.Println(b(1))
	fmt.Println(b(3))
	fmt.Println(b(4))

	// a := adder() is trivial and also works.
	a := adder2(0)
	var s1 int
	s1,a  = a(1)
	fmt.Println(s1,a)

	s1 ,a = a(2)
	fmt.Println(s1,a)
	//
	//fmt.Println(a(2))
	//for i := 0; i < 10; i++ {
	//	var s int
	//	s, a = a(i)
	//	fmt.Printf("0 + 1 + ... + %d = %d\n",
	//		i, s)
	//}
}
