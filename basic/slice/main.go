package main

import "fmt"

func feibo(n int){
	var slice = make([]int,n)
	slice[0] = 1
	slice[1] = 1
	for i:=2;i<n ;i++  {
		slice[i] = slice[i-1] +slice[i-2]
	}
	fmt.Println(slice)
}


func main()  {
	feibo(7)
}
