package main

import "fmt"

func isValid(s string) bool {
	n := len(s)
	if n % 2 == 1 {
		return false //判断奇数
	}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := []byte{}
	for i := 0; i < n; i++ {

		if pairs[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]//出栈
		} else {
			//如果不存在，说明是左括号， 入栈
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

func isVaild(str string) bool{

	pairs := map[byte]byte{
		'}' :'{',
		']' : '[',
		'>' :'<',
	}
	stack := []byte{}
	for i:= 0;i<len(str) ;i++  {
		//判断是否是右
		if pairs[str[i]] >0 {
			//右括号，
			//和堆顶对比,判断栈是否为空重要
			if len(stack) == 0 || stack[len(stack)-1] != pairs[str[i]] {
				return false
			}
			//对比成功出栈
			stack = stack[:len(stack)-1]

		}else{
			//左括号入栈
			stack = append(stack,str[i])
		}

	}
	return len(stack) == 0
}

func main(){
	str := "{}[]"
	fmt.Println(isVaild(str))
}
