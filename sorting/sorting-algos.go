// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main

import "fmt"

func main() {
	slice := []int{3, 1, 9, 2, 10, 2, 0, 2}
	// // {1,2,2,0,2,9,3,10}
	// pivotIndex := partition(slice)
	// fmt.Println("pivotIndex is", pivotIndex)
	// fmt.Println("partitioned slice is", slice)
	sortedArray := mergeSort(slice)
	fmt.Println("after sorting is", sortedArray)

}

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

func quickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	pivotIndex := partition(nums)

	quickSort(nums[0:pivotIndex])
	quickSort(nums[pivotIndex+1:])
}

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2

	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])

	return mergeSortedArray(left, right)
}

// merge two sorted array
func mergeSortedArray(nums1 []int, nums2 []int) []int {
	results := make([]int, 0, len(nums1)+len(nums2))
	i := 0
	j := 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			results = append(results, nums1[i])
			i++
		} else {
			results = append(results, nums2[j])
			j++
		}
	}
	for i < len(nums1) {
		results = append(results, nums1[i])
		i++
	}
	for j < len(nums2) {
		results = append(results, nums2[j])
		j++
	}
	return results

}
