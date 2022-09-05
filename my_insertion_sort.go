package my_sort

// package main

// import "fmt"

// func main() {
// 	arr := [5]int{1, 0, 5, 11, 9}
// 	myInsertionSort(arr[:])
// 	fmt.Println(arr)
// }

func myInsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		ind := i
		for ind > 0 && arr[ind-1] > arr[ind] {
			arr[ind-1], arr[ind] = arr[ind], arr[ind-1]
			ind--
		}
	}
}
