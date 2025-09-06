// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main

import "fmt"

func partition(nums []int) int {
	// pivot = last element
	fmt.Println("nums is", nums)
	// if pi
	i := 0
	pivot := nums[len(nums)-1]
	for index, _ := range nums {
		if nums[index] <= pivot {
			temp := nums[i]
			nums[i] = nums[index]
			nums[index] = temp
			i++
		}
	}
	fmt.Println("after partition nums is", nums)
	return i - 1

}

func main() {
	slice := []int{3, 1, 9, 2, 10, 2, 0, 2}
	// // {1,2,2,0,2,9,3,10}
	// pivotIndex := partition(slice)
	// fmt.Println("pivotIndex is", pivotIndex)
	// fmt.Println("partitioned slice is", slice)
	quickSort(slice)
	fmt.Println("after sorting is", slice)

}

func quickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	pivotIndex := partition(nums)

	quickSort(nums[0:pivotIndex])
	quickSort(nums[pivotIndex+1:])

}
