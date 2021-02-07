package main

import "fmt"

type Emp struct {
	id int
	name string
	next *Emp
}
type LinkNode struct{
	data *Emp
}
type Array struct {
	Link  [7]LinkNode
}

func (this *Array) hasFunc(id int)( postition int){
	return id % 7

}
func(this *LinkNode) Insert(emp *Emp){
	//
	fmt.Println("this:",this)
	if this.data == nil {
		this.data = emp
		fmt.Println("this.data",this.data)
		return
	}
	//不等于空的话，要比较大小
 	pre := this.data
	curr := this.data
	for  {
		if curr != nil {
			if curr.id > emp.id  {
				break;
			}
			pre = curr
			curr = curr.next

		}else{
			break
		}
	}
	pre.next = emp
	emp.next = curr


	return


}
func(this *Array) getById(id int)( emp Emp){
	return
}
func(this *Array) ListAll(){
	lenth := len(this.Link)
	for i:=0;i<lenth;i++ {
		fmt.Printf("数组 %d\n",i)
		//判断链表是否有数据
		this.Link[i].showLink()
	}
}
func(this LinkNode) showLink(){
	if this.data == nil {
		fmt.Printf("链表为空\n")
		return
	}

	curr := this.data
	for  {
		if curr != nil {
			fmt.Println("showlink",curr.id)
			curr = curr.next
		}else{
			break
		}
	}
}
func (this *Array)Insert(emp *Emp){
	no := this.hasFunc(emp.id)
	fmt.Println(no)
	this.Link[no].Insert(emp)
}

func main(){
	var arr Array
	emp1  := Emp{
		id:1,
		name:"西",
	}
	emp2  := Emp{
		id:8,
		name:"西2",
	}
	emp3  := Emp{
		id:15,
		name:"西3",
	}
	arr.Insert(&emp1)
	arr.Insert(&emp2)
	arr.Insert(&emp3)


	arr.ListAll()
}