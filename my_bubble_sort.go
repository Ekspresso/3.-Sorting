package my_sort

// package main

// import "fmt"

// func main() {
// 	arr := [5]int{1, 0, 5, 11, 9}
// 	myBubbleSort(arr[:])
// 	fmt.Println(arr)
// }

func myBubbleSort(arr []int) {
	swap := true
	for swap {
		swap = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swap = true
			}
		}
	}
}
