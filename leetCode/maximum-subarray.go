package main

import (
	"reflect"
	"fmt"
)
//53. 最大子序和
//给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//
//示例:
//
//输入: [-2,1,-3,4,-1,2,1,-5,4]
//输出: 6
//解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/maximum-subarray
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main(){
	nums :=[]int{-2,1,-3,4,-1,2,1,-5,6}
	a := reflect.TypeOf(nums)
	fmt.Println(a)
	//
	//max2 := maxTanxin(nums)
	//max3 := maxTanxin2(nums)
	//fmt.Println(max2)
	//fmt.Println(max3)
	//
	//
	//max := maxSubArray2(nums)//数组是传值引用？？？传递的数组会改变
	//fmt.Println(max)
	//fmt.Println(nums)

}
func maxSubArray3(nums []int) int {
	return get(nums, 0, len(nums) - 1).mSum;
}
func maxTanxin2(nums []int) int{
	if len(nums) ==0 {
		return -111111
	}
	cur_sum := nums[0]
	max_sum := cur_sum
	for i:=1;i<len(nums) ;i++  {
		cur_sum = max(nums[i],cur_sum+nums[i])
		max_sum = max(cur_sum,max_sum)

	}
	return max_sum

}

//贪心
func maxTanxin(nums []int) int{
	if len(nums) == 0 {
		return -1222222
	}
	var result int
	count := 0
	for i:=0;i< len(nums);i++  {

		count += nums[i]
		if count > result { // 取区间累计的最大值（相当于不断确定最大子序终止位置）
			result = count
		}
		if count <=0  {
			// 相当于重置最大子序起始位置，因为遇到负数一定是拉低总和
			count = 0
		}

	}
	return result
}

func pushUp(l, r Status) Status {
	iSum := l.iSum + r.iSum
	lSum := max(l.lSum, l.iSum + r.lSum)
	rSum := max(r.rSum, r.iSum + l.rSum)
	mSum := max(max(l.mSum, r.mSum), l.rSum + r.lSum)
	return Status{lSum, rSum, mSum, iSum}
}

func get(nums []int, l, r int) Status {
	if (l == r) {
		return Status{nums[l], nums[l], nums[l], nums[l]}
	}
	m := (l + r) >> 1
	lSub := get(nums, l, m)
	rSub := get(nums, m + 1, r)
	return pushUp(lSub, rSub)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type Status struct {
	lSum, rSum, mSum, iSum int
}

func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] + nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
func maxSubArray2(nums []int) int{
	max :=  nums[0]
	for i:= 1 ;i<len(nums) ;i++  {
		curr := nums[i]
		prew := nums[i-1]

		if  curr + prew > curr {
			curr = curr + prew
			nums[i] = curr
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
