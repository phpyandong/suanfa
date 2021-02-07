package main

import "fmt"

//
//func QuickSort2(nums []int,start,end int){
//	if start >= end{
//
//	}
//	position := QuickGetPosition(nums, start,end)
//	QuickSort2(nums,start,position-int(1))
//	QuickSort2(nums,start,position-1)
//
//}
func QuickSort(left int ,right int, array *[6]int){
	l:= left
	r:= right
	//pivot 是中轴 ，支点
	pivot := array[(left + right) /2]
	temp := 0
	// for 循环的目标
	for l<r {
		//从pivot 的左边找大于等于pivot的值
		for  array[l] < pivot {
			l ++
		}
		//从pivot 的右边找到小于等于pivot的值
		for  array[r] > pivot {
			r--
		}

		// l>=r 表明本次分解任务完成 ，brak
		if l >= r {
			break
		}

		//交换
		temp = array[l]
		array[l] = array[r]
		array[r] = temp
		////优化 相等的情况
		if array[l] == pivot {
			r--
		}
		if array[r] == pivot {
			l++
		}
	}
	////如果 l== r ,在移动下
	if l == r {
		l++
		r--
	}
	//向左递归
	if left < r {
		QuickSort(left, r, array)
	}

	//向右递归
	if right > l {
		QuickSort(l, right, array)
	}
}
//小到大
func bobosort(arr [6]int){
	finish := true
	for i:=0; i < len(arr) ;i++ {
		for j:=0; j<len(arr)-i -1; j++ {
			if(arr[j] > arr[j+1]){
				tmp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = tmp
				finish = false
			}
		}
		if finish == true {
			break
		}
	}
	fmt.Println(arr)

}
//选择
func choosesort(arr [6]int){
	for i:=0;i< len(arr) ;i++  {
		min := i
		for j:=i+1;j<len(arr) ;j++  {
			if arr[j]< arr[min] {
				min = j
			}
		}
		minValue := arr[min]
		arr[min] = arr[i]
		arr[i] = minValue
	}
	fmt.Println("select ",arr)
}

func inserSort(arr [6]int){
	for i:=1;i< len(arr) ; i++ {
		if arr[i] < arr[i-1] {
			tmp := arr[i]
			j := 0
			for j = i-1;j>=0&& tmp< arr[j] ;j--  {
				arr[j+1] = arr[j]
			}
			arr[j+1] = tmp
		}
	}
	fmt.Println("inset",arr)
}
func main(){
	arr := [6]int {-9,78,72,72,-567,68}
	fmt.Println("初始", arr)
	//调用快速排序
	//QuickSort(0, len(arr) - 1, &arr)
	inserSort(arr)
	fmt.Println("main..")
	fmt.Println(arr)
}