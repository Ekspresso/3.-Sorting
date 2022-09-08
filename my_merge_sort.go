package my_sort

// package main

// import "fmt"

// func main() {
// 	arr := [13]int{1, -99, 0, -86, 5, 11, 9, -5, -37, 121, 96, 525, 879}
// 	ret := myMergeSort(arr[:])
// 	fmt.Println(ret)
// }

// func myMergeSort(arr []int) []int {
// 	if len(arr) < 2 {
// 		return (arr)
// 	} else {
// 		left := 0
// 		right := len(arr) / 2
// 		larr := myMergeSort(arr[:right])
// 		rarr := myMergeSort(arr[right:])
// 		var sortedArr []int
// 		right = 0
// 		for i := 0; i < len(arr); i++ {
// 			if left < len(larr) && (right >= len(rarr) || larr[left] <= rarr[right]) {
// 				left = appendingArr(larr, &sortedArr, left)
// 			} else {
// 				right = appendingArr(rarr, &sortedArr, right)
// 			}
// 		}
// 		return (sortedArr)
// 	}
// }

// func appendingArr(arr []int, ret *[]int, i int) int {
// 	*ret = append(*ret, arr[i])
// 	for k := i + 1; k < len(arr); k++ {
// 		if arr[k-1] == arr[k] {
// 			*ret = append(*ret, arr[k])
// 		} else {
// 			return (k)
// 		}
// 	}
// 	return (len(arr))
// }

func myMergeSort(arr []int) []int {
	ret := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		ret[i] = arr[i]
	}
	if len(arr) > 1 {
		for !check(ret) {
			ret = merge(ret)
		}
		return ret
	} else {
		return ret
	}
}

// Функция создаёт новый массив и заполняет его поэлементно
// в порядке возрастания элементами из 2 частей исходного массива
func merge(arr []int) []int {
	left, right := cutArr(arr)
	ret := make([]int, len(left)+len(right))
	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			ret[i] = left[0]
			left = left[1:]
		} else {
			ret[i] = right[0]
			right = right[1:]
		}
		i++
	}
	for j := 0; j < len(left); j++ {
		ret[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		ret[i] = right[j]
		i++
	}
	return ret
}

// Функция возвращает результат разбиения массива на 2 подмассива
func cutArr(arr []int) ([]int, []int) {
	centre := len(arr) / 2
	left := make([]int, len(arr)-centre)
	right := make([]int, centre)
	j := 0
	for i := 0; i < len(arr); i++ {
		left[j] = arr[i]
		if i < len(arr)-1 {
			i++
			right[j] = arr[i]
		}
		j++
	}
	return left, right
}

func check(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}
