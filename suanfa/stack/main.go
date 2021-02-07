package main

import (
	"errors"
	"fmt"
)


type stack struct {
	MaxTop int	//表示我们栈最大可以存放数个数
	Top int		//表示栈顶，因为栈顶固定，因此我们直接使用Top
	arr [5]int  //数组模拟栈
}
//入栈
func (this *stack) Push(val int) (err error){
	//先判断栈是否满了
	if this.Top == this.MaxTop -1 {
		fmt.Println("stack full")
		return errors.New("stack full")
	}
	this.Top++
	this.arr[this.Top] = val
	return
}
func (this *stack) Pop()(val int ,err error){
	if this.Top == -1 {
		fmt.Println("stack empty!")
		return 0,errors.New("stack emtpy")
	}
	val = this.arr[this.Top]
	return val ,nil
}
func (this *stack) Lists(){
	if this.Top == -1 {
		fmt.Println("stack emtpy")
		return
	}
	for i:= this.Top;i >= 0;i--{
		fmt.Println("array【%d\n",i,this,this.arr[i])
	}
}
func main(){

}
