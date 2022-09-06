package my_sort

// package main

// import "fmt"

// func main() {
// 	arr := [5]int{1, 0, 5, 11, 9}
// 	myHeapSort(arr[:])
// 	fmt.Println(arr)
// }

func myHeapSort(arr []int) {
	// Первая сортировка кучи. Вытаскиваем max элемент на верх
	// Проход по дереву начинается с самого нижнего узла с дочерними элементами
	for i := len(arr)/2 - 1; i >= 0; i-- {
		maxHeapify(arr, len(arr), i)
	}
	// Цикл вытягивающий на верх кучи (arr[0]) максимальный элемент
	// и меняющий местами последний эл с первым, уменьшая размер кучи с конца
	for i := len(arr) - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		maxHeapify(arr, i, 0)
	}
}

// Функция сравнения элементов в узле. Если один из
// дочерних эл больше, то вытаскивает его в родительский
func maxHeapify(arr []int, heapSize int, ind int) {
	left := 2*ind + 1
	right := 2*ind + 2
	largest := ind
	if left < heapSize && arr[left] > arr[largest] {
		largest = left
	}
	if right < heapSize && arr[right] > arr[largest] {
		largest = right
	}
	if largest != ind {
		arr[ind], arr[largest] = arr[largest], arr[ind]
		maxHeapify(arr, heapSize, largest)
	}
}
