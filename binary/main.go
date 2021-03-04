package main

import "fmt"

func jinzhi(){
}
func main() {
	var a uint8 = 1
	var b uint8 = 2
	fmt.Println(a-b)
	//fmt.Printf("%b",uint8(1)-uint8(2))

	/**
		原码 = 最高位（符号位）+ 低位 （数值）
				符号 0正 1负
		例子 7 ：0000 0111
			-7：1000 0111
	 */
	fmt.Printf("%b",int8(-7))

	/**

		反码
			正数 = 原码
			负数  除 最高位 ，数值位 取反
		例子
			-7 : 1111 1000
	 */

	 /**
	 	补码
	 		正数 = 原码
	 		负数 反码 + 1
	 	例子
	   		-7 ： 1111 1000+1 = 1111 1001
	  */

	  /**
	  无符号 2^8位，有符号2^7
	  -1
	  原码 1000 0001
	  反码 1111 1110
	  补码 1111 1111

	   */

	//===========无符号整形================
	var ui8 uint8 = 255 //1111 1111  2^8 = 255
	fmt.Printf("unit:2^8 :%b\n",ui8)

	var ui16 uint16 = 65535  //1111 1111 1111 1111  2^16 = 65535
	fmt.Printf("unit16:2^16 :%b\n",ui16)
	var ui64 uint64 = 18446744073709551615
	//2^64 = 18446744073709551615
	//1111111111111111111111111111111111111111111111111111111111111111
	fmt.Printf("unit64:2^64 :%b\n",ui64)

	//========================================
	var i8 int8 = int8(127) // 111 1111  2^7 = 127
	fmt.Printf("int8 :2^7:%b\n",i8)

	var i16 int16 = int16(32767 )  // 111 1111 1111 1111 2^15 = 32767

	fmt.Printf("int16 :2^15:%b\n",i16)
	var i64 int64 = int64(18446744073709551615/2)   //2^63
	////2^64 = 18446744073709551615
	//1111111111111111111111111111111111111111111111111111111111111111
	fmt.Printf("int64 :2^63:%b\n",i64)
}