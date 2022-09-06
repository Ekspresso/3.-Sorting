package my_sort

func myCountingSort(arr []int) {
	shift := findMin(arr)
	k := findMax(arr) - shift
	counts := make([]int, k+1)
	// Считаем количество вхождений каждого числа в исходный массив
	for _, el := range arr {
		counts[el-shift] += 1
	}
	// Перезаписываем массив с кол-вами вхождений значениями
	// начальных индексов для каждого элемента
	startInd := 0
	for i := 0; i < k+1; i++ {
		count := counts[i]
		counts[i] = startInd
		startInd += count
	}
	// Заполнение нового массива отсортированными значениями
	sortedArr := make([]int, len(arr))
	for _, el := range arr {
		sortedArr[counts[el-shift]] = el
		counts[el-shift] += 1
	}
	// Вставка в исходный массив отсортированного
	for i := 0; i < len(arr); i++ {
		arr[i] = sortedArr[i]
	}
}

func findMin(arr []int) int {
	min := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return (min)
}

func findMax(arr []int) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return (max)
}
