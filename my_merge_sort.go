package my_sort

// package main

// import "fmt"

// func main() {
// 	arr := [13]int{1, -99, 0, -86, 5, 11, 9, -5, -37, 121, 96, 525, 879}
// 	ret := myMergeSort(arr[:])
// 	fmt.Println(ret)
// }

func myMergeSort(arr []int) []int {
	if len(arr) < 2 {
		return (arr)
	} else {
		left := 0
		right := len(arr) / 2
		larr := myMergeSort(arr[:right])
		rarr := myMergeSort(arr[right:])
		var sortedArr []int
		right = 0
		for i := 0; i < len(arr); i++ {
			if left < len(larr) && (right >= len(rarr) || larr[left] <= rarr[right]) {
				left = appendingArr(larr, &sortedArr, left)
			} else {
				right = appendingArr(rarr, &sortedArr, right)
			}
		}
		return (sortedArr)
	}
}

func appendingArr(arr []int, ret *[]int, i int) int {
	*ret = append(*ret, arr[i])
	for k := i + 1; k < len(arr); k++ {
		if arr[k-1] == arr[k] {
			*ret = append(*ret, arr[k])
		} else {
			return (k)
		}
	}
	return (len(arr))
}
