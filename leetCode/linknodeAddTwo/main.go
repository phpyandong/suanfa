package main

import "fmt"

type ListNode struct {

     Val int;
     Next *ListNode ;

}
func addTwoNumbers2(l1,l2 *ListNode) (listnode  *ListNode){
	var head *ListNode
	carry := 0//进位
	for l1 != nil || l2 != nil{
		n1,n2 := 0,0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1+n2+carry
		sum ,carry = (sum)%10,(sum)/10
		if head == nil {
			listnode = &ListNode{Val :sum}
			head = listnode
		}else{
			head.Next = &ListNode{sum,nil}
			head = head.Next
		}
	}
	if carry > 0 {
		head.Next = &ListNode{Val: carry}
	}
	return

}
func addTwoNumbers(l1, l2 *ListNode) (head *ListNode) {
	var tail *ListNode
	carry := 0 //进位
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return
}


func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil{
		return l2
	}
	if l2 == nil{
		return l1
	}
	var res *ListNode
	if l1.Val >= l2.Val{
		res = l2
		res.Next = mergeTwoLists(l1,l2.Next)
	}else{
		res = l1
		res.Next = mergeTwoLists(l1.Next,l2)
	}
	return res
}
func mergeTwo(l1 *ListNode,l2 *ListNode) (res *ListNode){
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val <= l2.Val {
		res = l1
		res.Next = mergeTwo(l1.Next,l2)
	}else{
		res = l2
		res.Next = mergeTwo(l1,l2.Next)
	}
	return res
}
func mergeDigui(l1,l2 *ListNode) (res *ListNode){
	if l1.Val >= l2.Val {
		res = l2
		res.Next = mergeDigui(l1,l2.Next)

	}else{
		res = l1
		res.Next = mergeDigui(l1.Next,l2)
	}
	return res
}
func merge(l1,l2 *ListNode )( res *ListNode){
	head := ListNode{}
	curr := &head
	for l1!= nil && l2 != nil {
		if l1.Val <= l2.Val {
			curr.Next = l1
			l1 = l1.Next
		}else{
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
	}
	if l2 != nil {
		curr.Next = l2
	}
	if l1 != nil {
		curr.Next = l1
	}
	next := head.Next
	return next
}
//输入：(1 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：6 -> 0 -> 8
//原因：342 + 465 = 807
func main(){
	node1_1 := ListNode{1,nil}
	node1_2 := ListNode{2,nil}
	node1_3 := ListNode{9,nil}
	node1_2.Next = &node1_3
	node1_1.Next = &node1_2
	node2_1 := ListNode{4,nil}
	node2_2 := ListNode{5,nil}
	node2_3 := ListNode{6,nil}
	node2_2.Next = &node2_3
	node2_1.Next = &node2_2
	newnode  := addTwoNumbers2(&node1_1, &node2_1)
	newnode = mergeTwo(&node1_1,&node2_1)

	for newnode != nil {
		fmt.Println(newnode.Val)
		newnode = newnode.Next
	}

}
