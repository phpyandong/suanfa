package main

import "fmt"

type HeroNode struct{
	no int
	name string
	nickname string
	pre *HeroNode
	next *HeroNode
}
func InsertHeroNode(head *HeroNode,newHeroNode *HeroNode){
	temp := head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	temp.next = newHeroNode
	newHeroNode.pre = temp
}
func InsertHeroNode2(head *HeroNode,newHeroNode *HeroNode){
	temp := head
	flag := true
	for{
		if temp.next  == nil{
			break
		}else if temp.next.no >= newHeroNode.no{
			break
		}else if temp.next.no == newHeroNode.no{
			flag = false
			break
		}
		temp = temp.next
	}
	if !flag {
		fmt.Println("对不起，已经存在no=",newHeroNode.no)
		return
	}else{
		newHeroNode.next = temp.next
		newHeroNode.pre = temp
		if temp.next !=nil{
			temp.next.pre = newHeroNode
		}
		temp.next = newHeroNode
	}
}

func DelHeroNode(head *HeroNode,id int){
	temp := head
	flag := false
	for{
		if temp.next ==nil {
			break
		}else if temp.next.no == id{
			flag = true
			break
		}
		temp = temp.next

	}
	if flag {
		temp.next = temp.next.next
		if temp.next != nil {
			temp.next.pre = temp
		}
	}else{
		fmt.Println("不存在")
	}

}
func ListHeroNode(head *HeroNode){
	fmt.Println("")

	fmt.Println("打印开始")
	temp := head
	if temp.next == nil{
		return
	}
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	for{
		fmt.Printf("[%d,%s,%s] ===>",temp.no,temp.name,temp.nickname)
		temp = temp.pre
		if temp.pre == nil {
			break
		}
	}
	fmt.Println("")

	fmt.Println("打印结束")
}
func main(){
	head := &HeroNode{}
	hero1 := &HeroNode{
		no:1,
		name:"松江",
		nickname:"及时雨",
	}
	hero2 := &HeroNode{
		no:2,
		name :"卢俊义",
		nickname:"玉麒麟",
	}
	hero3 := &HeroNode{
		no:2,
		name :"林冲",
		nickname :"豹子头",
	}
	hero4 := &HeroNode{
		no :3,
		name :"无用",
		nickname:"智多星",
	}

	InsertHeroNode2(head,hero3)
	InsertHeroNode2(head,hero1)
	InsertHeroNode2(head,hero2)
	InsertHeroNode2(head,hero4)
	ListHeroNode(head)
	DelHeroNode(head,2)
	ListHeroNode(head)
}