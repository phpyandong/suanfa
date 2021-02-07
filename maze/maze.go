package main

import (
	"fmt"
	"os"
)

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])

		}
	}

	return maze
}

type point struct {
	i, j int
}

var dirs = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}
//判断当前点是否合法，是否越界，合法的话得到值
func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}

	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}

	fmt.Println("gridij:",grid[p.i][p.j])

	res := 	 grid[p.i][p.j]

	fmt.Println(grid)
	return res, true
}
func walk2 (maze [][]int, start ,end point) [][]int{
	//构造路径
	steps := make([][]int, len(maze))

	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	Queue := []point{start}
	for len(Queue) > 0 {
		//弹出第一个点
		cur := Queue[0]
		Queue = Queue[1:]
		if cur == end {
			break //发现终点退出
		}
		//上下左右四个方向探索
		for _, dir := range dirs {
			next := cur.add(dir)
			val, ok := next.at(maze)//探索增加下一个位置
			if !ok || val == 1 {// 不ok 或者 撞墙
				continue
			}

			//判断当前点是否合法
			curStep, _ := cur.at(steps)
			steps[next.i][next.j] =
				curStep + 1
			Queue = append(Queue, next)//把下一个点加入队列

		}

	}
	return steps

}
func walk3(maze [][]int,start,end point) [][]int{
	//增加一个map 记录走过的位置
	step := make([][]int,len(maze))//和原map一致
	//for i:=0;i<len(step) ;i++  {
	//	step[i] = make([]int ,len(maze[i]))
	//}
	for i := range step  {
		step[i] = make([]int,len(maze[i]))
	}

	//创建一个queue 存储下一个位置,	//将起点入列
	Queues := []point{start}
	//Queues = Queues[1:] 不报错

	for len(Queues) > 0 {
		//探索一个方向
		curr := Queues[0]
		Queues = Queues[1:]

		if curr == end {
			break //发现终点退出
		}

		addPoint := point{1, 0}
		step,Queues = AddPoints(curr,addPoint,&maze,start,step ,Queues )
		fmt.Println("我是外部的que:",Queues)
		////向下得到point
		//nextPoint := curr.add(addPoint)
		//
		////将next point 加入到queue
		////取出next 的合法性
		//NextValueFromMap,ok := nextPoint.at(maze)
		////如果
		//if !ok || NextValueFromMap == 1 {
		//	continue
		//	//撞墙
		//}
		//if nextPoint == start {
		//	continue   //回到起点不合法
		//}
		//NextValuesFromStep,ok := nextPoint.at(step)
		//if !ok || NextValuesFromStep != 0 {
		//	//位置信息已经改变了，说明走过了
		//	continue//不合法
		//
		//}
		//CurrValue ,ok := curr.at(step)
		//fmt.Println("接收curr",CurrValue)
		//
		//step[nextPoint.i][nextPoint.j] = CurrValue+1
		//
		//
		//Queues = append(Queues, nextPoint )
		//fmt.Println("最佳后 ",Queues)


	}
 	return step
}
func AddPoints(curr point,addPoint point,maze [][]int,start point,step [][]int,Queues []point)( stepss [][]int,Queuess []point){
	//向下得到point
	nextPoint := curr.add(addPoint)

	//将next point 加入到queue
	//取出next 的合法性
	NextValueFromMap,ok := nextPoint.at(maze)
	//如果
	if !ok || NextValueFromMap == 1 {
		return
		//撞墙
	}
	if nextPoint == start {
		return    //回到起点不合法
	}
	NextValuesFromStep,ok := nextPoint.at(step)
	if !ok || NextValuesFromStep != 0 {
		//位置信息已经改变了，说明走过了
		return //不合法

	}
	CurrValue ,ok := curr.at(step)
	fmt.Println("接收curr",CurrValue)

	step[nextPoint.i][nextPoint.j] = CurrValue+1
	stepss = step

	Queues = append(Queues, nextPoint )
	Queuess = Queues
	fmt.Println("最佳后 ",Queues)
	return stepss,Queuess
}

func walk(maze [][]int,
	start, end point) [][]int {
	//构造路径
	steps := make([][]int, len(maze))

	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}



	Queue := []point{start}

	for len(Queue) > 0 {
		//弹出第一个点
		cur := Queue[0]
		Queue = Queue[1:]

		if cur == end {
			break //发现终点退出
		}
		//上下左右四个方向探索
		for _, dir := range dirs {
			next := cur.add(dir)
			fmt.Println("next:",next)
			val, ok := next.at(maze)//探索增加下一个位置
			if !ok || val == 1 {// 不ok 或者 撞墙
				continue
			}
			//因为用了两个map 记录原数据和新数据
			val, ok = next.at(steps)//判断下一个位置是否合法，越界
			if !ok || val != 0 {//不ok 或者 已经走过了，防止后退
				continue
			}

			if next == start {//不可以回起点
				continue
			}
			//判断当前点是否合法
			curStep, _ := cur.at(steps)//仅仅是获取当前curr的value
			steps[next.i][next.j] =
				curStep + 1

			Queue = append(Queue, next)//把下一个点加入队列
			fmt.Println(Queue)
		}


	}

	return steps
}

func main() {
	maze := readMaze("maze/maze.in")

	//steps := walk3(maze, point{0, 0},
	//	point{len(maze) - 1, len(maze[0]) - 1})

	steps := walk3(maze, point{0, 0},
		point{4, 0})

	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}

	// TODO: construct path from steps
}
