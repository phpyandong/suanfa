package main

import "fmt"
type User struct {
	Test []int
	Name string
}
func main(){
	m := map[string]User{"aa":{[]int{1},"张三"}}
	m["aa"].Test[0] = 1
	fmt.Println(m)
	//m["aa"].Name = "李四"
	fmt.Println(m["aa"].Name)
}