package main

import "fmt"
func SetWay2(myMap *[8][7]int,i int,j int) bool{
	if  myMap[5][1] == 2{
		fmt.Println("退出的条件，是i j 到达坐标 5，1")
	}
	if myMap[i][j] == 0 {//执行条件
		myMap[i][j] = 2
		if SetWay(myMap,i+1,j) {//逼近条件
			return true
		}else{
			myMap[i][j] = 3//死路
		}

	}else{
		//不可执行，
		return false
	}
	return true
}

//编写一个函数，完成老鼠找路
//myMap *[8][7]int:地图，保证是同一个地图，使用引用 //i,j 表示对地图的哪个点进行测试
func SetWay(myMap *[8][7]int, i int, j int) bool {
	//分析出什么情况下，就找到出路 //myMap[6][5] == 2

	if myMap[5][1] == 2 {
		fmt.Println("找到出路了")
		return true
	} else { //说明要继续找
		if myMap[i][j] == 0 { //如果这个点是可以探测
			//假设这个点是可以通, 但是需要探测 上下左右 //换一个策略 下右上左
			myMap[i][j] = 2
			fmt.Println(i,"i")
			fmt.Println(j,"j")

			if SetWay(myMap, i+1, j) { //下
				fmt.Println("下")
				return true
			}else{
				fmt.Println("else")
				return false
			}
			//else{
			//	myMap[i][j] = 3
			//	return false
			//}
			//else { // 死路
			//	fmt.Println("死路")
			//	myMap[i][j] = 3
			//	return false
			//}
		} else { // 说明这个点不能探测，为 1，是强
		 	return false
		}
	}
}

func main() {
	//先创建一个二维数组，模拟迷宫
	//规则
	//1. 如果元素的值为 1 ，就是墙
	//2. 如果元素的值为 0, 是没有走过的点
	//3. 如果元素的值为 2, 是一个通路
	//4. 如果元素的值为 3， 是走过的点，但是走不通
	var myMap [8][7]int
	//先把地图的最上和最下设置为 1
	for i := 0; i < 7; i++ {
		myMap[0][i] = 1
		myMap[7][i] = 1
		//先把地图的最左和最右设置为 1 for i := 0 ; i < 8 ; i++ {
		myMap[i][0] = 1
		myMap[i][6] = 1
	}
	//myMap[3][1] = 0
	//myMap[3][2] = 0
	//myMap[2][1] = 1
	//myMap[2][2] = 1
	//myMap[2][3] = 1
	//myMap[2][4] = 1


	//myMap[5][1] = 2
	//myMap[2][2] =
	//输出地图
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(myMap[i][j], " ")
		}
		fmt.Println()
	}
	//使用测试
	SetWay(&myMap, 1, 1)
	fmt.Println("探测完毕....") //输出地图
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(myMap[i][j], " ")
		}
		fmt.Println()
	}
}
