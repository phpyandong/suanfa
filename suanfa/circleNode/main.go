package main

import "fmt"

type CatNode struct {
	no int
	name string
	next *CatNode
}
func InsertCatNode(head *CatNode,newCatNode *CatNode){
	if head.next == nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head//构成一个环形
		fmt.Println(newCatNode,"加入到环形的链表")
		return
	}
	//定义一个临时变量，帮忙，找到环形的最后节点

	temp := head

	for{
		if temp.next == head{
			break
		}
		temp = temp.next
	}
	temp.next = newCatNode
	newCatNode.next = head

}
func ListCircleLink(head *CatNode){
	fmt.Println("环形链表的情况如下：")
	temp := head
	if temp.next == nil {
		fmt.Println("空空的")
		return
	}
	for{
		fmt.Printf("猫的信息为=[id=%d name=%s] ->\n", temp.no, temp.name)
		if temp.next == head{
			break
		}
		temp = temp.next
	}
}
func DelCatNode(head *CatNode,id int) *CatNode{
	temp := head
	helper := head
	if temp.next == nil{
		fmt.Println("空的不可删除")
		return head
	}
	//如果只有一个节点
	if temp.next == head{
		if temp.no == id{
			temp.next = nil
		}
		return head
	}
	//将helper 定位到链表最后
	for{
		if helper.next == head{
			break
		}
		helper = helper.next
	}
	//如果有两个包含两个以上的节点

	flag := true
	for{
		if temp.next == head {
			break
		}
		if temp.no == id {
			//删除的是头结点，这个需要特别注意，结合示意图
			if temp == head{
				head = head.next
			}
			//恭喜找到，我们也可以在直接删除
			helper.next = temp.next
			fmt.Println("删除了猫咪%d\n",id)
			flag = false
			break
		}
		temp = temp.next//移动
		helper = helper.next//移动
	}
	if flag {
		if temp.no == id{
			helper.next = temp.next
			fmt.Printf("删除猫猫=%d\n", id)
		}else{
			fmt.Println("对不起没有%d",id)
		}
	}
	return head

}
func main(){

	//这里我们初始化一个环形链表的头结点
	head := &CatNode{}
	//创建一只猫
	cat1 := &CatNode{
		no : 1,
		name : "tom", }
	cat2 := &CatNode{ no : 2,
		name : "tom2", }
	cat3 := &CatNode{ no : 3,
		name : "tom3", }
	InsertCatNode(head, cat1)
	InsertCatNode(head, cat2)
	InsertCatNode(head, cat3)
	ListCircleLink(head)
	head = DelCatNode(head, 1)
	fmt.Println()
	fmt.Println()
	fmt.Println()
	ListCircleLink(head)
}