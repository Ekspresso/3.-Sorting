package my_sort

// package main

// import "fmt"

// func main() {
// 	arr := [5]int{1, 0, 5, 11, 9}
// 	myQuickSort(arr[:], 0, len(arr)-1)
// 	fmt.Println(arr)
// }

func myQuickSort(arr []int, left int, right int) {
	if left < right {
		pivot := partition(arr, left, right)
		myQuickSort(arr, left, pivot-1)
		myQuickSort(arr, pivot+1, right)
	}
}

func partition(arr []int, left int, right int) int {
	pivot := arr[right]
	i := left - 1
	for j := left; j < right; j++ {
		if arr[j] <= pivot {
			i++
			arr[j], arr[i] = arr[i], arr[j]
		}
	}
	arr[right], arr[i+1] = arr[i+1], arr[right]
	return (i + 1)
}
